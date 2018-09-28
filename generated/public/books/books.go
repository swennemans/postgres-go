// Code generated by gnorm, DO NOT EDIT!

package books

import (
	"context"
	"database/sql"
	"time"

	"github.com/Mattel/mcpp/data/gnorm"
	"github.com/gnormal/postgres-go/generated"
	"github.com/gnormal/postgres-go/generated/public/enum"
	uuid "github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// TableName is the primary table that this particular gnormed file deals with.
const TableName = "books"

// Row represents a row from 'books'.
type Row struct {
	ID        int            // id (PK)
	AuthorID  uuid.UUID      // author_id
	Available time.Time      // available
	Booktype  enum.BookType  // booktype
	Isbn      string         // isbn
	Pages     int            // pages
	Summary   sql.NullString // summary
	Title     string         // title
}

// Field values for every column in Books.
var (
	AuthorIDCol  generated.UuidUUIDField      = "author_id"
	AvailableCol generated.TimeTimeField      = "available"
	BooktypeCol  enum.BookTypeField           = "booktype"
	IDCol        generated.IntField           = "id"
	IsbnCol      generated.StringField        = "isbn"
	PagesCol     generated.IntField           = "pages"
	SummaryCol   generated.SqlNullStringField = "summary"
	TitleCol     generated.StringField        = "title"
)

// All retrieves all rows from 'books' as a slice of Row.
func All(ctx context.Context, db generated.DB) ([]*Row, error) {
	const sqlstr = `SELECT
		author_id, available, booktype, id, isbn, pages, summary, title
		FROM public.books`

	var vals []*Row
	q, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, errors.Wrap(err, "query Books")
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.AuthorID,
			&r.Available,
			&r.Booktype,
			&r.ID,
			&r.Isbn,
			&r.Pages,
			&r.Summary,
			&r.Title,
		)
		if err != nil {
			return nil, errors.Wrap(err, "all Books")
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// CountQuery retrieve one row from 'books'.
func CountQuery(ctx context.Context, db gnorm.DB, where gnorm.WhereClause) (int, error) {
	const origsqlstr = `SELECT
		count(*) as count
		FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	count := 0
	err := db.QueryRowContext(ctx, sqlstr, where.Values()...).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "count Books")
	}
	return count, nil
}

// Query retrieves rows from 'books' as a slice of Row.
func Query(ctx context.Context, db generated.DB, where generated.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT
		id, author_id, isbn, booktype, title, pages, summary, available
		FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.QueryContext(ctx, sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.Wrap(err, "query Books")
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID,
			&r.AuthorID,
			&r.Isbn,
			&r.Booktype,
			&r.Title,
			&r.Pages,
			&r.Summary,
			&r.Available,
		)
		if err != nil {
			return nil, errors.Wrap(err, "query Books")
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// QueryOrder retrieves rows from 'books' as a slice of Row in a particular order.
func QueryOrder(ctx context.Context, db generated.DB, where generated.WhereClause, orderby generated.OrderBy) ([]*Row, error) {
	const origsqlstr = `SELECT
		id, author_id, isbn, booktype, title, pages, summary, available
		FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + orderby.String()

	var vals []*Row
	q, err := db.QueryContext(ctx, sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.Wrap(err, "query Books")
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID,
			&r.AuthorID,
			&r.Isbn,
			&r.Booktype,
			&r.Title,
			&r.Pages,
			&r.Summary,
			&r.Available,
		)
		if err != nil {
			return nil, errors.Wrap(err, "query Books")
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// One retrieve one row from 'books'.
func One(ctx context.Context, db generated.DB, where generated.WhereClause) (*Row, error) {
	const origsqlstr = `SELECT
		id, author_id, isbn, booktype, title, pages, summary, available
		FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	r := &Row{}
	err := db.QueryRowContext(ctx, sqlstr, where.Values()...).Scan(&r.ID,
		&r.AuthorID,
		&r.Isbn,
		&r.Booktype,
		&r.Title,
		&r.Pages,
		&r.Summary,
		&r.Available,
	)
	if err != nil {
		return nil, errors.Wrap(err, "queryOne Books")
	}
	return r, nil
}

// First retrieve one row from 'books' when sorted by orderby.
func First(ctx context.Context, db generated.DB, where generated.WhereClause, orderby generated.OrderBy) (*Row, error) {
	const origsqlstr = `SELECT
		id, author_id, isbn, booktype, title, pages, summary, available
		FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + orderby.String()

	r := &Row{}
	err := db.QueryRowContext(ctx, sqlstr, where.Values()...).Scan(&r.ID,
		&r.AuthorID,
		&r.Isbn,
		&r.Booktype,
		&r.Title,
		&r.Pages,
		&r.Summary,
		&r.Available,
	)
	if err != nil {
		return nil, errors.Wrap(err, "queryFirst Books")
	}
	return r, nil
}

// Find retrieves a row from 'books' by its primary key(s).
func Find(ctx context.Context, db generated.DB,
	id int,
) (*Row, error) {
	const sqlstr = `SELECT
		author_id, available, booktype, id, isbn, pages, summary, title
	FROM public.books WHERE ( id = $1 )`

	r := &Row{}
	err := db.QueryRowContext(ctx, sqlstr,
		id,
	).Scan(&r.AuthorID,
		&r.Available,
		&r.Booktype,
		&r.ID,
		&r.Isbn,
		&r.Pages,
		&r.Summary,
		&r.Title,
	)
	if err != nil {
		return nil, errors.Wrap(err, "find Books")
	}
	return r, nil
}

// Insert inserts the row into the database.
func Insert(ctx context.Context, db generated.DB, r *Row) error {
	const sqlstr = `INSERT INTO public.books ` +
		`(
			author_id, available, booktype, isbn, pages, summary, title
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		) ` +
		`RETURNING
			created_at, id
		`

	err := db.QueryRowContext(ctx, sqlstr, &r.AuthorID,

		&r.Available,

		&r.Booktype,

		&r.Isbn,

		&r.Pages,

		&r.Summary,

		&r.Title,
	).Scan(&r.CreatedAt, &r.ID)

	return errors.Wrap(err, "insert Books")
}

// InsertIgnore inserts the row into the database but ignores conflicts
func InsertIgnore(ctx context.Context, db gnorm.DB, r *Row, constraint string) error {
	sqlstr := `INSERT INTO public.books ` +
		`(
			author_id, available, booktype, isbn, pages, summary, title
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		) ` +
		`ON CONFLICT ON CONSTRAINT ` + constraint + ` DO NOTHING `
	_, err := db.ExecContext(ctx, sqlstr, r.AuthorID, r.Available, r.Booktype, r.Isbn, r.Pages, r.Summary, r.Title)
	return errors.Wrap(err, "insert ignore Books")
}

// Set sets a single column on an existing row in the database.
func Set(ctx context.Context, db generated.DB, set generated.Where, where generated.WhereClause) (int64, error) {
	idx := 2
	sqlstr := `UPDATE public.books SET ` +
		set.Field + " = $1 " +
		` WHERE ` +
		where.String(&idx)

	res, err := db.ExecContext(ctx, sqlstr, append([]interface{}{set.Value}, where.Values()...)...)
	if err != nil {
		return 0, errors.Wrap(err, "set Books")
	}
	return res.RowsAffected()
}

// AppendInt64 adds a value to a field
func AppendInt64(ctx context.Context, db generated.DB, name string, value interface{}, where generated.WhereClause) (int64, error) {
	idx := 2
	sqlstr := `UPDATE public.books SET ` +
		name + " = array_append(" + name + ", $1::bigint) " +
		` WHERE ` +
		where.String(&idx)

	res, err := db.ExecContext(ctx, sqlstr, append([]interface{}{value}, where.Values()...)...)
	if err != nil {
		return 0, errors.Wrap(err, "append_int64 Books")
	}
	return res.RowsAffected()
}

// Inc increments the value of a single column on an existing row in the database.
func Inc(ctx context.Context, db generated.DB, inc generated.Where, where generated.WhereClause) (int64, error) {
	idx := 2
	sqlstr := `UPDATE public.books SET ` +
		inc.Field + " = " + inc.Field + " + $1" +
		` WHERE ` +
		where.String(&idx)

	res, err := db.ExecContext(ctx, sqlstr, append([]interface{}{inc.Value}, where.Values()...)...)
	if err != nil {
		return 0, errors.Wrap(err, "inc Books")
	}
	return res.RowsAffected()
}

// Upsert performs an insert-or-update in one DB call for Books.
// Unlike insert, upsert requires that you have set any IDs on the row you're upserting.
// NOTE: PostgreSQL 9.5+ only
func Upsert(ctx context.Context, db generated.DB, r *Row) error {

	const sqlstr = `INSERT INTO public.books (
		author_id, available, booktype, id, isbn, pages, summary, title
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	) ON CONFLICT (id) DO UPDATE SET (
		author_id, available, booktype, id, isbn, pages, summary, title
	) = (
		$1, $2, $3, $4, $5, $6, $7, $8
	)`

	_, err := db.ExecContext(ctx, sqlstr, r.AuthorID, r.Available, r.Booktype, r.ID, r.Isbn, r.Pages, r.Summary, r.Title)
	return errors.Wrap(err, "upsert Books")
}

// Delete deletes the Row from the database. Returns the number of items deleted.
func Delete(ctx context.Context,
	db generated.DB,
	id int,
) (int64, error) {
	const sqlstr = `DELETE FROM public.books 
	WHERE
	  id = $1
	`

	res, err := db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return 0, errors.Wrap(err, "delete Books")
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// DeleteWhere deletes Rows from the database and returns the number of rows deleted.
func DeleteWhere(ctx context.Context, db generated.DB, where generated.WhereClause) (int64, error) {
	const origsqlstr = `DELETE FROM public.books WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	res, err := db.ExecContext(ctx, sqlstr, where.Values()...)
	if err != nil {
		return 0, errors.Wrap(err, "delete Books")
	}
	return res.RowsAffected()
}

// DeleteAll deletes all Rows from the database and returns the number of rows deleted.
func DeleteAll(ctx context.Context, db generated.DB) (int64, error) {
	const sqlstr = `DELETE FROM public.books`

	res, err := db.ExecContext(ctx, sqlstr)
	if err != nil {
		return 0, errors.Wrap(err, "deleteall Books")
	}
	return res.RowsAffected()
}