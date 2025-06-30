package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"social-network/config"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(dbPath string) error {
	_, filename, _, _ := runtime.Caller(0)
	baseDir, _ := strings.CutSuffix(filename, "backend/pkg/db/sqlite/sqlite.go")
	DBPath := baseDir + config.DBPath + "/" + config.DBName
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	defer CloseDatabase(db)

	if err := ApplyMigrations(db); err != nil {
		return err
	}
	log.Println("Database migrations applied successfully")
	return nil
}

func ApplyMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}
	_, filename, _, _ := runtime.Caller(0)
	baseDir, _ := strings.CutSuffix(filename, "backend/pkg/db/sqlite/sqlite.go")
	migrationPath := "file://" + baseDir + config.MigPath

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"sqlite3",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration initialization failed: %w", err)
	}
	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	return nil
}

func CloseDatabase(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Failed to close database: %v", err)
	}
}
