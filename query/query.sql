-- name: CreateBook :one
INSERT INTO books (title, author, year)
VALUES (?, ?, ?)
RETURNING *;
