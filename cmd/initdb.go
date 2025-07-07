package cmd

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

//go:embed schema.sql
var ddl string

var initDBCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Initialize the database schema",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := sql.Open("sqlite3", "books.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer conn.Close()

		if _, err := conn.ExecContext(context.Background(), ddl); err != nil {
			log.Fatalf("Failed to initialize schema: %v", err)
		}

		fmt.Println("Database initialized successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initDBCmd)
}
