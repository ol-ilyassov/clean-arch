package contact

import (
	"context"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"

	"github.com/jackc/pgx/v4"
)

type Contact interface {
	CreateContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error)
}
