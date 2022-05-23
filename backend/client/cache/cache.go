package cache

import (
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
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
	// SelectAllCardsWithFilter generate a select query for all stored cards with a given filter.
	SelectAllCardsWithFilter(filter model.CardFilter) (string, error)
	// InsertCard generates a query to insert a specific card into the database.
	InsertCard(card *model.Card) (string, error)
}

// YgoCache handles the caching of generic data in a postgres database.
type YgoCache struct {
	Client         DatabaseClient
	QueryTemplater QueryGenerator
}

// NewYgoCache creates a new cache by the given name.
func NewYgoCache(client DatabaseClient) (*YgoCache, error) {
	logrus.Debug("Cache -> Create new cache")

	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &YgoCache{
		Client:         client,
		QueryTemplater: queryTemplater,
	}, nil
}
