package postgres

import (
	"ol-ilyassov/clean_arch/pkg/tools/transaction"
	"ol-ilyassov/clean_arch/pkg/type/context"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/repository/contact/postgres/dao"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) CreateContact(c context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.CreateContactTx(ctx, tx, contacts...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) CreateContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	if len(contacts) == 0 {
		return []*contact.Contact{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"slurm", "contact"},
		dao.CreateColumnContact,
		r.toCopyFromSource(contacts...))
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *Repository) UpdateContact(c context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContact(c context.Context, ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListContact(c context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) ReadContactByID(c context.Context, ID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (r *Repository) CountContact(ctx context.Context) (uint64, error) {
	panic("implement me")
}
