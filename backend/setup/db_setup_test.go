package setup_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ygodraft/backend/client/cache/mocks"
	"ygodraft/backend/setup"
)

func TestTestNewDatabase(t *testing.T) {
	t.Run("create a new setup", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}

		// when
		databaseSetup := setup.NewDatabaseSetup(dbMock)

		// then
		assert.NotNil(t, databaseSetup)
	})
}
