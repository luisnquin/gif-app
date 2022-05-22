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
)

const (
	Postgres DBMS = 1
	MySQL    DBMS = 2
)

var (
	ErrFailedToSaveInDB = errors.New("failed to save in database")
)

var DB Database

type (
	Database struct {
		db *sql.DB
	}

	DBMS int
)

/*
	type Databaser interface {
		Query(query string, args ...any) (*sql.Rows, error)
		QueryRow(query string, args ...any) *sql.Row
		Exec(query string, args ...any) (sql.Result, error)
	}
*/

func Connect(dbms DBMS) func() error {
	var (
		driver string
		dsn    string
	)

	switch dbms {
	case Postgres:
		if utils.IsRunningInADockerContainer() {
			dsn = fmt.Sprintf(config.Server.Database.InContainerDSN, os.Getenv("HOST"))
		} else {
			dsn = config.Server.Database.InLocalDSN
		}

		driver = "postgres"

	case MySQL:
	}

	database, err := sql.Open(driver, dsn)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err = database.Ping(); err != nil {
		panic("no response from database: " + err.Error())
	}

	DB.db = database

	return DB.db.Close
}

func (d *Database) QueryRow(query string, args ...any) *sql.Row {
	timeOut := time.Second * time.Duration(config.Server.Database.SecondsToTimeOut)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) Query(query string, args ...any) (*sql.Rows, error) {
	timeOut := time.Second * time.Duration(config.Server.Database.SecondsToTimeOut)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return rows, nil
}

func (d *Database) Exec(query string, args ...any) (sql.Result, error) {
	timeOut := time.Second * time.Duration(config.Server.Database.SecondsToTimeOut)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql query failed: %w", err)
	}

	return result, nil
}
