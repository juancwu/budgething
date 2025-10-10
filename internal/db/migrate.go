package db

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"

	"github.com/pressly/goose/v3"
)

var dialectMap = map[string]string{
	"sqlite": "sqlite3",
	"pgx":    "postgres",
}

func getDialect(driver string) string {
	dialect, ok := dialectMap[driver]
	if ok {
		return dialect
	}
	return driver
}

func setupGoose(driver string) error {
	err := goose.SetDialect(getDialect(driver))
	if err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	migrationsDir, err := fs.Sub(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create sub-filesystem migrations directory: %w", err)
	}

	goose.SetBaseFS(migrationsDir)

	return nil
}

func RunMigrations(db *sql.DB, driver string) error {
	err := setupGoose(driver)
	if err != nil {
		return err
	}

	err = goose.Up(db, ".")
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	slog.Info("migrations completed successfully")

	return nil
}
