package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"ol-ilyassov/clean_arch/pkg/tools/transaction"
	"ol-ilyassov/clean_arch/pkg/type/columnCode"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/repository/storage/postgres/dao"
)

// What fields could be used for sorting list.
var mappingSortContact = map[columnCode.ColumnCode]string{
	"id":          "id",
	"fullName":    "full_name",
	"phoneNumber": "phone_number",
	"name":        "name",
	"surname":     "surname",
	"patronymic":  "patronymic",
	"email":       "email",
	"gender":      "gender",
	"age":         "age",
}

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx) // transaction start.
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err) // transaction end declaration.
	}(ctx, tx) // our written Finish function decides Commit or Rollback on the base of error.

	// main creation function.
	response, err := r.createContactTx(ctx, tx, contacts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Repository) createContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	// no contacts to create.
	if len(contacts) == 0 {
		return []*contact.Contact{}, nil
	}

	// CopyFrom function is good for large number of data to import(as example), or create.
	// Difficulty case, make check, is some contact with same name, surname exists(?).
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"slurm", "contact"},
		dao.CreateColumnContact,
		r.toCopyFromSource(contacts...),
	)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *Repository) UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	upContact, err := r.oneContactTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}
	in, err := updateFn(upContact)
	if err != nil {
		return nil, err
	}

	return r.updateContactTx(ctx, tx, in)
}

func (r *Repository) updateContactTx(ctx context.Context, tx pgx.Tx, in *contact.Contact) (*contact.Contact, error) {

	builder := r.genSQL.Update("slurm.contact").
		Set("email", in.Email().String()).
		Set("phone_number", in.PhoneNumber().String()).
		Set("age", in.Age()).
		Set("gender", in.Gender()).
		Set("modified_at", in.ModifiedAt()).
		Set("name", in.Name().String()).
		Set("surname", in.Surname().String()).
		Set("patronymic", in.Patronymic().String()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":          in.ID(),
				"is_archived": false,
			},
		}).
		Suffix(`RETURNING
			id,
			created_at,
			modified_at,
			phone_number,
			email,
			name,
			surname,
			patronymic,
			age,
			gender`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoContacts []*dao.Contact
	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return nil, err
	}

	return r.toDomainContact(daoContacts[0])
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteContactTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteContactTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	builder := r.genSQL.Update("slurm.contact").
		Set("is_archived", true).
		Set("modified_at", time.Now().UTC()).
		Where(squirrel.Eq{"is_archived": false, "id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	var daoContacts []*dao.Contact
	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return err
	}

	// check is contact exists in group and delete it from them.
	if err = r.updateGroupsContactCountByFilters(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if parameter.Pagination.Limit == 0 {
		parameter.Pagination.Limit = r.options.DefaultLimit
	}

	contacts, err := r.listContactTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *Repository) listContactTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"modified_at",
		"phone_number",
		"email",
		"name",
		"surname",
		"patronymic",
		"age",
		"gender",
	).From("slurm.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		// parsing for converting values to string for SQL query.
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortContact)...)
	} else {
		builder = builder.OrderBy("created_at DESC")
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(parameter.Pagination.Limit)
	}
	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(parameter.Pagination.Offset)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoContacts []*dao.Contact
	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return nil, err
	}

	// converter.go => Convert to Domain representation.
	return r.toDomainContacts(daoContacts)
}

// there could be the use of filters.
// there could be a wrapper of listContactTx with Limit 1.
// So by basic written oneContactTx is duplicate for listContactTx, first one could be erased.
func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err = r.oneContactTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) oneContactTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (*contact.Contact, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"modified_at",
		"phone_number",
		"email",
		"name",
		"surname",
		"patronymic",
		"age",
		"gender",
	).From("slurm.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false, "id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoContact []*dao.Contact
	if err = pgxscan.ScanAll(&daoContact, rows); err != nil {
		return nil, err
	}

	if len(daoContact) == 0 {
		return nil, errors.New("contact not found")
	}

	return r.toDomainContact(daoContact[0])
}

func (r *Repository) CountContact() (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From("slurm.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var row = r.db.QueryRow(context.Background(), query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
