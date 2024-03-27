package model

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// DatabaseClient is used to communicate against a database.
type DatabaseClient interface {
	// QueryRow performs a query and return a single row.
	QueryRow(query string, args ...any) (pgx.Row, error)
	// Query perform a query and returns a pgx.Rows containing multiple rows.
	Query(query string, args ...any) (pgx.Rows, error)
	// Exec executes a query against the database.
	Exec(query string, args ...any) (pgconn.CommandTag, error)
	// Select executes a query against the database and scans the result into the targetObject.
	Select(query string, targetObject any, args ...any) error
}
