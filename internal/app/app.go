package app

import (
	"fmt"

	"git.juancwu.dev/juancwu/budgething/internal/config"
	"git.juancwu.dev/juancwu/budgething/internal/db"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Cfg *config.Config
	DB  *sqlx.DB
}

func New(cfg *config.Config) (*App, error) {
	database, err := db.Init(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	err = db.RunMigrations(database.DB, cfg.DBDriver)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return &App{
		Cfg: cfg,
		DB:  database,
	}, nil
}

func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}
