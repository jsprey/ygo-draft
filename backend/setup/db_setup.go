package setup

import (
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/client/cache"
)

//go:embed queries/cards.sql
var createTableCards string

//go:embed queries/card_sets.sql
var createTableCardSets string

//go:embed queries/users.sql
var createTableUsers string

// DatabaseSetup is responsible to setup the database including the creation of the database and the data tables.
type DatabaseSetup struct {
	Client cache.DatabaseClient
}

func NewDatabaseSetup(client cache.DatabaseClient) *DatabaseSetup {
	return &DatabaseSetup{Client: client}
}

func (ds *DatabaseSetup) Setup() error {
	logrus.Debug("Setup -> Database -> Creating `cards` table")
	_, err := ds.Client.Exec(createTableCards)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	logrus.Debug("Setup -> Database -> Creating `card_sets` table")
	_, err = ds.Client.Exec(createTableCardSets)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	logrus.Debug("Setup -> Database -> Creating `users` table")
	_, err = ds.Client.Exec(createTableUsers)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}
