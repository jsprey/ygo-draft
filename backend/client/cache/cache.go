package cache

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type DatabaseClient interface {
	QueryRow(query string, args ...any) (pgx.Row, error)
	Query(query string, args ...any) (pgx.Rows, error)
	Exec(query string, args ...any) (pgconn.CommandTag, error)
	Select(query string, targetObject any, args ...any) error
}

// YgoCache handles the caching of generic data in a postgres database.
type YgoCache struct {
	Client DatabaseClient
}

// NewYgoCache creates a new cache by the given name.
func NewYgoCache(client DatabaseClient) (*YgoCache, error) {
	logrus.Debug("Cache -> Create new cache")

	cache := &YgoCache{
		Client: client,
	}

	return cache, nil
}
