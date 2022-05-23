package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	// PostgreSQL driver.
	_ "github.com/lib/pq"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/utils"
	"go.uber.org/fx"
)

// const (
// 	Postgres DBMS = 1
// 	MySQL    DBMS = 2
// )

// type DBMS int

var ErrFailedToSaveInDB = errors.New("failed to save in DB")

type database struct {
	config *config.Configuration
	db     *sql.DB
}

type Databaser interface {
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func New(lc fx.Lifecycle, config *config.Configuration) Databaser {
	var dsn string

	if utils.IsRunningInADockerContainer() {
		dsn = fmt.Sprintf(config.Database.InContainerDSN, os.Getenv("HOST"))
	} else {
		dsn = config.Database.InLocalDSN
	}

	var DB database

	DB.config = config

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			database, err := sql.Open("postgres", dsn)
			if err != nil {
				return fmt.Errorf("failed to connect to database: %w", err)
			}

			if err = database.Ping(); err != nil {
				return fmt.Errorf("no response from database: %w", err)
			}

			DB.db = database

			return nil
		},
		OnStop: func(ctx context.Context) error {
			DB.db.Close()

			return nil
		},
	})

	return &DB
}

func (d *database) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	timeOut := time.Second * d.config.Database.SecondsToTimeOut

	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *database) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	timeOut := time.Second * d.config.Database.SecondsToTimeOut

	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return rows, nil
}

func (d *database) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	timeOut := time.Second * d.config.Database.SecondsToTimeOut

	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return result, nil
}
