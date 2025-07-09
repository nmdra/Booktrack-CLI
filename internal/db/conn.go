package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const dbPath = "db.sqlite"

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return conn, nil
}
