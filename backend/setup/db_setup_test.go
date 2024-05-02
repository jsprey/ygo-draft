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
		dbMock := mocks.NewDatabaseClient(t)
		usermgtMock := &mocks.UsermgtClient{}

		// when
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		// then
		assert.NotNil(t, databaseSetup)
	})
}

func TestDatabaseSetup_Setup(t *testing.T) {
	t.Run("fail as retrieving admin throws error", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		usermgtMock := &mocks.UsermgtClient{}
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetUser", config.AdminUserEmail).Return(nil, assert.AnError)

		// when
		err = databaseSetup.Setup()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock, usermgtMock)
	})

	t.Run("perform setup fails on error", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		usermgtMock := &mocks.UsermgtClient{}
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, assert.AnError)

		// when
		err = databaseSetup.Setup()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})

	t.Run("creates admin account if not existent", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetUser", config.AdminUserEmail).Return(nil, model.ErrorUserDoesNotExist.WithParam("admin@admin"))
		usermgtMock.On("CreateUser", mock.Anything).Return(nil)
		dbMock.On("Select", "SELECT d.id, d.display_name, d.description, d.rarity, d.tags, d.costs\nFROM ygodraft.public.draft_store AS d", mock.Anything).Return(nil)

		// when
		err = databaseSetup.Setup()

		// then
		require.NoError(t, err)
	})

	t.Run("fail setup when getting error on retrieving products catalog", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetUser", config.AdminUserEmail).Return(nil, nil)
		dbMock.On("Select", "SELECT d.id, d.display_name, d.description, d.rarity, d.tags, d.costs\nFROM ygodraft.public.draft_store AS d", mock.Anything).Return(assert.AnError)

		// when
		err = databaseSetup.Setup()

		// then
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("create products catalog when not existent", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetUser", config.AdminUserEmail).Return(nil, nil)
		dbMock.On("Select", "SELECT d.id, d.display_name, d.description, d.rarity, d.tags, d.costs\nFROM ygodraft.public.draft_store AS d", mock.Anything).Return(nil)

		// when
		err = databaseSetup.Setup()

		// then
		require.NoError(t, err)
	})

	t.Run("run setup without creating anything new", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		defer mock.AssertExpectationsForObjects(t, dbMock)
		usermgtMock := &mocks.UsermgtClient{}
		defer mock.AssertExpectationsForObjects(t, usermgtMock)
		databaseSetup, err := setup.NewDatabaseSetup(dbMock, usermgtMock)
		require.NoError(t, err)

		dbMock.On("Exec", mock.Anything).Return(nil, nil)
		usermgtMock.On("GetUser", config.AdminUserEmail).Return(nil, nil)

		expectedProduct1 := model.DraftStoreProduct{
			ID: "new", DisplayName: "New Product", Description: "This is new", Rarity: model.ProductRarityCommon, Costs: 500,
		}
		expectedProduct1.SetTags(model.ProductTagLoseCard)
		expectedProduct2 := model.DraftStoreProduct{
			ID: "new2", DisplayName: "New Product2", Description: "This is new 2", Rarity: model.ProductRarityCommon, Costs: 500,
		}
		expectedProduct2.SetTags(model.ProductTagGainCard)
		expectedProducts := []*model.DraftStoreProduct{&expectedProduct1, &expectedProduct2}

		dbMock.On("Select", "SELECT d.id, d.display_name, d.description, d.rarity, d.tags, d.costs\nFROM ygodraft.public.draft_store AS d", mock.Anything).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			expectedCards, ok := arg1.(*[]*model.DraftStoreProduct)
			if !ok {
				t.FailNow()
			}

			*expectedCards = expectedProducts
		}).Return(nil)

		// when
		err = databaseSetup.Setup()

		// then
		require.NoError(t, err)
	})
}
