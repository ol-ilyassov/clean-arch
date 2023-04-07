package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Group interface {
	UpdateGroupsContactCountByFilters(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error
}
