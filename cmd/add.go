/*
Copyright © 2025 NIMENDRA <nimendraonline@gmail.com>
*/
package cmd

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nmdra/Booktrack-CLI/db"
	"github.com/spf13/cobra"
)

var (
	title  string
	author string
	year   int32
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new book",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := sql.Open("sqlite3", "books.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer conn.Close()

		q := db.New(conn)

		book, err := q.CreateBook(context.Background(), db.CreateBookParams{
			Title:  title,
			Author: author,
			Year: sql.NullInt64{
				Int64: int64(year),
				Valid: year != 0,
			},
		})
		if err != nil {
			log.Fatalf("Failed to create book: %v", err)
		}

		fmt.Printf("Book added: %s by %s (ID %d)\n", book.Title, book.Author, book.ID)
	},
}

func init() {
	addCmd.Flags().StringVar(&title, "title", "", "Book title")
	addCmd.Flags().StringVar(&author, "author", "", "Book author")
	addCmd.Flags().Int32Var(&year, "year", 0, "Published year")

	addCmd.MarkFlagRequired("title")
	addCmd.MarkFlagRequired("author")

	rootCmd.AddCommand(addCmd)
}
