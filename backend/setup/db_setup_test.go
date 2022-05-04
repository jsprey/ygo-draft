package setup_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
func TestDatabaseSetup_Setup(t *testing.T) {
	t.Run("perform setup", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		databaseSetup := setup.NewDatabaseSetup(dbMock)
		dbMock.On("Exec", mock.Anything).Return(nil, nil)

		// when
		err := databaseSetup.Setup()

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, dbMock)
	})

	t.Run("perform setup fails on error", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		databaseSetup := setup.NewDatabaseSetup(dbMock)
		dbMock.On("Exec", mock.Anything).Return(nil, assert.AnError)

		// when
		err := databaseSetup.Setup()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
}
