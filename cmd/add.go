/*
Copyright © 2025 NIMENDRA <nimendraonline@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

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
			if isNoSuchTableError(err) {
				fmt.Fprintf(os.Stderr, "Database not initialized. Run `booktrack-cli init` to set it up.\n")
			} else {
				fmt.Fprintf(os.Stderr, "Failed to add book: %v\n", err)
			}
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

func isNoSuchTableError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "no such table")
}
