package cache

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

// YgoCache handles the caching of generic data in a postgres database.
type YgoCache struct {
	Client         model.DatabaseClient
	QueryTemplater model.YgoQueryGenerator
}

// NewYgoCache creates a new cache by the given name.
func NewYgoCache(client model.DatabaseClient) (*YgoCache, error) {
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
