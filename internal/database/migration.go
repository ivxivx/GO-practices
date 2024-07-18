package database

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required for migration with postgres
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*
var MigrationFiles embed.FS

func Migrate(dbURL string) (uint, uint, error) {
	driver, err := iofs.New(MigrationFiles, "migrations")
	if err != nil {
		return 0, 0, fmt.Errorf("could not create fs: %w", err)
	}

	migrateInstance, errM := migrate.NewWithSourceInstance("iofs", driver, dbURL)
	if errM != nil {
		return 0, 0, fmt.Errorf("could not migrate conn string: %w", errM)
	}

	before, dirty, errV := migrateInstance.Version()
	if errV != nil && !errors.Is(errV, migrate.ErrNilVersion) {
		return 0, 0, fmt.Errorf("could not migrate conn string: %w", errV)
	}

	if dirty {
		return 0, 0, fmt.Errorf("could not migrate conn string: %w", errV)
	}

	if err := migrateInstance.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return 0, 0, fmt.Errorf("could not migrate up: %w", err)
	}

	after, dirty, errV := migrateInstance.Version()
	if errV != nil {
		return 0, 0, fmt.Errorf("could not migrate conn string: %w", errV)
	}

	if dirty {
		return 0, 0, fmt.Errorf("could not migrate conn string: %w", errV)
	}

	return before, after, nil
}
