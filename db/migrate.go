package db

import (
	"database/sql"
	"io/fs"

	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

func MigrateUp(database *sql.DB, logger logger.Logger, files fs.FS) {
	driver, _ := postgres.WithInstance(database, &postgres.Config{})

	dir, err := iofs.New(files, "migrations")
	if err != nil {
		logger.Error("migration up failed: %v", err)
		return
	}

	m, _ := migrate.NewWithInstance(
		"iofs", dir, "postgres", driver)

	err = m.Up()
	if err != nil {
		logger.Error("migration up failed: %v", err)
		return
	}

	logger.Info("Database migration - up finished")
}

func MigrateDown(database *sql.DB, logger logger.Logger, files fs.FS) {
	driver, _ := postgres.WithInstance(database, &postgres.Config{})

	dir, err := iofs.New(files, "migrations")
	if err != nil {
		logger.Error("migration down failed: %v", err)
		return
	}

	m, _ := migrate.NewWithInstance(
		"iofs", dir, "postgres", driver)

	err = m.Down()
	if err != nil {
		logger.Error("migration down failed: %v", err)
		return
	}

	logger.Info("Database migration - down finished")
}

func MigrateFixUp(database *sql.DB, logger logger.Logger, files fs.FS, ver int) {
	driver, _ := postgres.WithInstance(database, &postgres.Config{})

	dir, err := iofs.New(files, "migrations")
	if err != nil {
		logger.Error("migration fix failed: %v", err)
		return
	}

	m, _ := migrate.NewWithInstance(
		"iofs", dir, "postgres", driver)

	err = m.Force(ver)
	if err != nil {
		logger.Error("migration force failed: %v", err)
		return
	}

	err = m.Up()
	if err != nil {
		logger.Error("migration up failed: %v", err)
		return
	}

	logger.Info("Database migration - fix finished")
}

func MigrateFixDown(database *sql.DB, logger logger.Logger, files fs.FS, ver int) {
	driver, _ := postgres.WithInstance(database, &postgres.Config{})

	dir, err := iofs.New(files, "migrations")
	if err != nil {
		logger.Error("migration fixdown failed: %v", err)
		return
	}

	m, _ := migrate.NewWithInstance(
		"iofs", dir, "postgres", driver)

	err = m.Force(ver)
	if err != nil {
		logger.Error("migration force failed: %v", err)
		return
	}

	err = m.Down()
	if err != nil {
		logger.Error("migration down failed: %v", err)
		return
	}

	logger.Info("Database migration - fixdown finished")
}