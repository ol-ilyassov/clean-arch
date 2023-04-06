// get connection Pool to Postgres DB
// Use approach throw ENV (could be changed).
package postgres

// used to init some DB connection.

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/spf13/viper/remote"
)

// the use of
// https://github.com/pkg/errors
// will add stack of errors - provision of context and place of error.

// generally unnecessary code, but interesting.
// configurations could be directly taken from config places (Consul, config files, etc...).
func init() {
	// Could be used another variant. Viper, ...
	// if dev launched as prod (mb remote configs), then initDefaultEnv() (env variables) could be unused.
	// Consul will directly write into structure Settings.
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

// Default values - use case: Development mode, containers in Docker, Local machines
// Env to pg params (default env params for pgxpool lib):
// "PGHOST":               "host",
// "PGPORT":               "port",
// "PGDATABASE":           "database",
// "PGPASSWORD":           "password",
// "PGPASSFILE":           "passfile",
// "PGAPPNAME":            "application_name",
// "PGCONNECT_TIMEOUT":    "connect_timeout",
// "PGSSLMODE":            "sslmode",
// "PGSSLKEY":             "sslkey",
// "PGSSLCERT":            "sslcert",
// "PGSSLROOTCERT":        "sslrootcert",
// "PGTARGETSESSIONATTRS": "target_session_attrs",
// "PGSERVICE":            "service",
// "PGSERVICEFILE":        "servicefile".

// Needed, when your approach has default values, an user didn't enter config parameters and ENV parameters.
// Default values wouldn't be comparable with production values.
// Be default, production values will be taken from another source (config file, or centralized config manager).
func initDefaultEnv() error {
	// read from env variables.
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "localhost"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "5432"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "test"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGUSER", "test"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "test_123456"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "disable"); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

type Store struct {
	Pool *pgxpool.Pool
}

// used when the approach with external config is in work (from Consul, etc...).
type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func (s Settings) toDSN() string {
	var args []string
	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}

	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}

	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", s.Database))
	}

	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}

	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}

	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}
	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(settings.toDSN()) // - could be empty string, if ENV variant used.
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// context could be used with timer ...
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Pool: conn}, nil
}
