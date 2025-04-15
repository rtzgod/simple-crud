package repository

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgres(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to postgres"))
	}
	err = db.Ping()
	if err != nil {
		panic(errors.Wrap(err, "failed to check connection with postgres"))
	}
	if err := startMigration(db.DB); err != nil {
		return nil, err
	}
	return db, nil
}

func startMigration(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to create database driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"postgres", driver)
	if err != nil {
		return errors.Wrap(err, "failed to create migration instance")
	}
	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return errors.Wrap(err, "failed to get migration version")
	}

	log.Printf("Current migration version: %d, dirty: %t", version, dirty)

	if version == 0 {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return errors.Wrap(err, "failed to run migration")
		}
	}
	return nil
}
