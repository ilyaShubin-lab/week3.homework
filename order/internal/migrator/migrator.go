package migrator

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

// подключение и путь к .sql-файлам
type Migrator struct {
	db            *sql.DB
	migrationsDir string
}

func NewMigrator(db *sql.DB, migrationsDir string) *Migrator {
	return &Migrator{
		db:            db,
		migrationsDir: migrationsDir,
	}
}

func (m *Migrator) Up() error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	err = goose.Up(m.db, m.migrationsDir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}
	return nil
}
