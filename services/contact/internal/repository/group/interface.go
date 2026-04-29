package group

import (
	"ol-ilyassov/clean_arch/pkg/type/context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Group interface {
	UpdateGroupsContactCountByFilters(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error
}
