package store

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	// PostgreSQL driver.
	_ "github.com/lib/pq"
	"github.com/luisnquin/meow-app/src/server/config"
)

const (
	Postgres DBMS = 1
	MySQL    DBMS = 2
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
		password := os.Getenv("POSTGRES_PASSWORD")
		database := os.Getenv("POSTGRES_DB")
		user := os.Getenv("POSTGRES_USER")
		port := os.Getenv("DB_PORT")
		host := os.Getenv("HOST")

		dsn = "dbname=%s user=%s password=%s host=%s port=%s sslmode=disable"
		dsn = fmt.Sprintf(dsn, database, user, password, host, port)

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