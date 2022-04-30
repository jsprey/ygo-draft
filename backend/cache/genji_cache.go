package cache

import (
	"context"
	"fmt"
	"github.com/genjidb/genji"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	QueryCreateTable = "CREATE TABLE %s"
)

// YgoCache handles the caching of generic data.
type YgoCache struct {
	GenjiDB *genji.DB
}

func isGenjinAlreadyExistsError(err error) bool {
	return strings.Contains(err.Error(), "already exists")
}

// NewYgoCache creates a new cache by the given name
func NewYgoCache(cacheName string) (*YgoCache, error) {
	logrus.Debugf("Cache -> Create new cache [%s]", cacheName)

	db, err := genji.Open(cacheName)
	if err != nil {
		return nil, fmt.Errorf("failed to open internal database: %w", err)
	}
	db = db.WithContext(context.Background())

	cache := &YgoCache{
		GenjiDB: db,
	}

	err = cache.createTables()
	if err != nil {
		return nil, fmt.Errorf("failed to create tables for database: %w", err)
	}

	return cache, nil
}

func (yc *YgoCache) createTables() error {
	logrus.Debug("Cache -> Creating all tables")

	err := yc.createCardsTable()

	if isGenjinAlreadyExistsError(err) {
		logrus.Debugf("Cache -> Table [%s] already exists -> skip creation", err.Error())
	} else if err != nil {
		return err
	}

	return nil
}

func (yc *YgoCache) createTable(tableName string) error {
	query := fmt.Sprintf(QueryCreateTable, tableName)

	logrus.Debugf("Cache -> Creating table with query [%s]", query)

	err := yc.GenjiDB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to exec query [%s]: %w", query, err)
	}

	return nil
}

// Close closes the connection to the internal database.
func (yc *YgoCache) Close() error {
	return yc.GenjiDB.Close()
}
