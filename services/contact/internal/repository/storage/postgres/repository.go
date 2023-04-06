package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

// repository implements the interface of useCase/adapters/storage/interface.go

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options Options
}

type Options struct{} // optional, example: request timer

func New(db *pgxpool.Pool, o Options) *Repository {
	var r = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(o)
	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		// example why if is needed. To make log only if data is changed.
		r.options = options
	}
}
