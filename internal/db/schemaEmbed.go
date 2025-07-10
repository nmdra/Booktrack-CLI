package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "modernc.org/sqlite"
)

const dbFile = "db.sqlite"

//go:embed schema.sql
var schemaFile embed.FS

func InitDatabase() error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}
	defer db.Close()
	
	// Read the SQL file content
	sqlBytes, err := schemaFile.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read embedded schema: %w", err)
	}

	if _, err := db.ExecContext(context.Background(), string(sqlBytes)); err != nil {
		return fmt.Errorf("failed to run schema: %w", err)
	}

	fmt.Println("Database created.")

	return nil
}
