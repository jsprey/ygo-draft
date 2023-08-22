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

// QueryGenerator is used to generate the queries send to the database client.
type QueryGenerator interface {
	// SelectCardByID generate a select query for a specific card by the given id.
	SelectCardByID(id int) (string, error)
	// SelectAllCards generate a select query for all stored cards.
	SelectAllCards() (string, error)
	// SelectAllSets generate a select query for all stored sets.
	SelectAllSets() (string, error)
	// SelectSetByCode generate a select query for a given stored set.
	SelectSetByCode(cardSet string) (string, error)
	// SelectAllCardsWithFilter generate a select query for all stored cards with a given filter.
	SelectAllCardsWithFilter(filter CardFilter) (string, error)
	// InsertCard generates a query to insert a specific card into the database.
	InsertCard(card *Card) (string, error)
	// InsertSet generates a query to insert a specific card set into the database.
	InsertSet(set CardSet) (string, error)
}
