package book

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	appdb "github.com/nmdra/Booktrack-CLI/internal/db"
	"github.com/nmdra/Booktrack-CLI/internal/repository"
)

type Book struct {
	Title  string
	Author string
	Year   int32
}

func Add(input Book) error {
	conn, err := appdb.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	q := repository.New(conn)

	book, err := q.CreateBook(context.Background(), repository.CreateBookParams{
		Title:  input.Title,
		Author: input.Author,
		Year:   sqlNullInt(input.Year),
	})
	if err != nil {
		return fmt.Errorf("create book: %w", err)
	}

	slog.Info("Book added",
		"title", book.Title,
		"author", book.Author,
		"id", book.ID,
	)

	return nil
}

func sqlNullInt(v int32) sql.NullInt64 {
	if v == 0 {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: int64(v), Valid: true}
}
