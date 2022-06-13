package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	// PostgreSQL driver.
	_ "github.com/lib/pq"
	"github.com/luisnquin/gif-app/src/server/config"
	"github.com/luisnquin/gif-app/src/server/utils"
)

var ErrFailedToSaveInDB = errors.New("failed to save in DB")

func initPostgresClient(config *config.Configuration) (*sql.DB, error) {
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

	return db, nil
}
