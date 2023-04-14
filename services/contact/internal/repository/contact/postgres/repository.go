package postgres

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db     *pgxpool.Pool
	genSQL squirrel.StatementBuilderType

	options Options
}

type Options struct {
	Timeout       time.Duration
	DefaultLimit  uint64
	DefaultOffset uint64
}

func New(db *pgxpool.Pool, o Options) (*Repository, error) {

	var r = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(o)
	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}
	if options.Timeout == 0 {
		options.Timeout = time.Second * 30
	}

	if r.options != options {
		r.options = options
	}
}
