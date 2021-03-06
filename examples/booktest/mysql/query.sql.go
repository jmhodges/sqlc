// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package booktest

import (
	"context"
	"time"
)

const booksByTitleYear = `-- name: BooksByTitleYear :many
select book_id, author_id, isbn, book_type, title, yr, available, tags from books where title = ? and yr = ?
`

type BooksByTitleYearParams struct {
	Title string
	Yr    int
}

type BooksByTitleYearRow struct {
	BookID    int
	AuthorID  int
	Isbn      string
	BookType  BookTypeType
	Title     string
	Yr        int
	Available time.Time
	Tags      string
}

func (q *Queries) BooksByTitleYear(ctx context.Context, arg BooksByTitleYearParams) ([]BooksByTitleYearRow, error) {
	rows, err := q.db.QueryContext(ctx, booksByTitleYear, arg.Title, arg.Yr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BooksByTitleYearRow
	for rows.Next() {
		var i BooksByTitleYearRow
		if err := rows.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Yr,
			&i.Available,
			&i.Tags,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createAuthor = `-- name: CreateAuthor :exec
insert into authors(name) values (?)
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createAuthor, name)
	return err
}

const createBook = `-- name: CreateBook :exec
insert into books(author_id, isbn, booktype, title, yr, available, tags) values (?, ?, ?, ?, ?, ?, ?)
`

type CreateBookParams struct {
	AuthorID  int
	Isbn      string
	Unknown   interface{}
	Title     string
	Yr        int
	Available time.Time
	Tags      string
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) error {
	_, err := q.db.ExecContext(ctx, createBook,
		arg.AuthorID,
		arg.Isbn,
		arg.Unknown,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tags,
	)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
delete from books where book_id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, book_id int) error {
	_, err := q.db.ExecContext(ctx, deleteBook, book_id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
select author_id, name from authors where author_id = ?
`

type GetAuthorRow struct {
	AuthorID int
	Name     string
}

func (q *Queries) GetAuthor(ctx context.Context, author_id int) (GetAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, author_id)
	var i GetAuthorRow
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const getBook = `-- name: GetBook :one
select book_id, author_id, isbn, book_type, title, yr, available, tags from books where book_id = ?
`

type GetBookRow struct {
	BookID    int
	AuthorID  int
	Isbn      string
	BookType  BookTypeType
	Title     string
	Yr        int
	Available time.Time
	Tags      string
}

func (q *Queries) GetBook(ctx context.Context, book_id int) (GetBookRow, error) {
	row := q.db.QueryRowContext(ctx, getBook, book_id)
	var i GetBookRow
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tags,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :exec
update books set title = ?, tags = ? where book_id = ?
`

type UpdateBookParams struct {
	Title  string
	Tags   string
	BookID int
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook, arg.Title, arg.Tags, arg.BookID)
	return err
}

const updateBookISBN = `-- name: UpdateBookISBN :exec
update books set title = ?, tags = ?, isbn = ? where book_id = ?
`

type UpdateBookISBNParams struct {
	Title  string
	Tags   string
	Isbn   string
	BookID int
}

func (q *Queries) UpdateBookISBN(ctx context.Context, arg UpdateBookISBNParams) error {
	_, err := q.db.ExecContext(ctx, updateBookISBN,
		arg.Title,
		arg.Tags,
		arg.Isbn,
		arg.BookID,
	)
	return err
}
