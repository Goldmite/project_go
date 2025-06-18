package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./internal/database/bookshelf.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	fmt.Println("Connected to the SQLite database successfully.")

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	sqlBytes, err := os.ReadFile("internal/database/schema.sql")
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("failed to read db schema: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to execute schema: %w", err)
	}

	return db, nil
}
