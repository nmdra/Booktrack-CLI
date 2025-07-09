/*
Copyright © 2025 NIMENDRA <nimendraonline@gmail.com>
*/
package cmd

import (
	_ "embed"
	"log/slog"
	"os"

	"github.com/nmdra/Booktrack-CLI/internal/book"
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
		err := book.Add(book.Book{
			Title:  title,
			Author: author,
			Year:   year,
		})
		if err != nil {
			slog.Error("Failed to add book", "error", err)
			os.Exit(1)
		}
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
