package postgres

import (
	"fmt"
	"ol-ilyassov/clean_arch/services/contact/internal/repository/contact"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MIGRATIONS_DIR", "./services/contact/internal/repository/storage/postgres/migrations")
}

type Repository struct {
	db     *pgxpool.Pool
	genSQL squirrel.StatementBuilderType

	repoContact contact.Contact

	options Options
}

type Options struct {
	Timeout       time.Duration
	DefaultLimit  uint64
	DefaultOffset uint64
}

func New(db *pgxpool.Pool, repoContact contact.Contact, o Options) (*Repository, error) {
	if err := migrations(db); err != nil {
		return nil, err
	}

	var r = &Repository{
		genSQL:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		repoContact: repoContact,
		db:          db,
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

func migrations(pool *pgxpool.Pool) (err error) {
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		return err
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			err = errClose
			return
		}
	}()

	dir := viper.GetString("MIGRATIONS_DIR")
	goose.SetTableName("contact_version")
	if err = goose.Run("up", db, dir); err != nil {
		return fmt.Errorf("goose %s error : %w", "up", err)
	}
	return
}
