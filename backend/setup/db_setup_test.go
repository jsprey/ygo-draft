package setup_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/model/mocks"
	"ygodraft/backend/setup"
)

func TestTestNewDatabase(t *testing.T) {
	t.Run("create a new setup", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		usermgtMock := &mocks.UsermgtClient{}

		// when
		databaseSetup := setup.NewDatabaseSetup(dbMock, usermgtMock)

		// then
		assert.NotNil(t, databaseSetup)
	})
}

func TestDatabaseSetup_Setup(t *testing.T) {
	t.Run("fail as retrieving admin throws error", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup := setup.NewDatabaseSetup(dbMock, usermgtMock)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetCurrentUser", config.AdminUserEmail).Return(nil, assert.AnError)

		// when
		err := databaseSetup.Setup()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("perform setup fails on error", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		usermgtMock := &mocks.UsermgtClient{}
		databaseSetup := setup.NewDatabaseSetup(dbMock, usermgtMock)

		dbMock.On("Exec", mock.Anything).Return(nil, assert.AnError)

		// when
		err := databaseSetup.Setup()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})

	t.Run("creates admin account if not existent", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup := setup.NewDatabaseSetup(dbMock, usermgtMock)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetCurrentUser", config.AdminUserEmail).Return(nil, model.ErrorUserDoesNotExist.WithParam("admin@admin"))
		usermgtMock.On("CreateUser", mock.Anything).Return(nil)

		// when
		err := databaseSetup.Setup()

		// then
		require.NoError(t, err)
	})

	t.Run("run setup without creatin admin account as it already exists", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup := setup.NewDatabaseSetup(dbMock, usermgtMock)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetCurrentUser", config.AdminUserEmail).Return(nil, nil)

		// when
		err := databaseSetup.Setup()

		// then
		require.NoError(t, err)
	})
}
