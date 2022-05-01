package cache

import (
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type PostgresClient interface {
	QueryRow(query string, args ...any) pgx.Row
	Query(query string, args ...any) (pgx.Rows, error)
	Exec(query string, args ...any) error
	Select(dst interface{}, query string, args ...interface{}) error
}

// YgoCache handles the caching of generic data in a postgres database.
type YgoCache struct {
	Client PostgresClient
}

// NewYgoCache creates a new cache by the given name.
func NewYgoCache(client PostgresClient) (*YgoCache, error) {
	logrus.Debug("Cache -> Create new cache")

	cache := &YgoCache{
		Client: client,
	}

	return cache, nil
}
