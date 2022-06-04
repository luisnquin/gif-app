package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	// PostgreSQL driver.
	_ "github.com/lib/pq"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/utils"
)

var ErrFailedToSaveInDB = errors.New("failed to save in DB")

func initPostgresClient(config *config.Configuration) (Querier, error) {
	var dsn string

	if utils.IsRunningInADockerContainer() {
		dsn = config.Database.InContainerDSN
	} else {
		dsn = config.Database.InLocalDSN
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &database{
		config: config,
		db:     db,
	}, nil
}

type database struct {
	config *config.Configuration
	db     *sql.DB
}

type Querier interface {
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Ping() error
}

func (d *database) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	ctx, cancel := context.WithTimeout(ctx, d.config.Database.SecondsToTimeOut*time.Second)
	defer cancel()

	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *database) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(ctx, d.config.Database.SecondsToTimeOut*time.Second)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return rows, nil
}

func (d *database) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(ctx, d.config.Database.SecondsToTimeOut*time.Second)
	defer cancel()

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return result, nil
}

func (d *database) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return d.db.PingContext(ctx)
}
