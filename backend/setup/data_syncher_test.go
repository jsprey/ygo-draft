package setup_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/client/ygo/mocks"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/setup"
)

func TestNewYgoDataSyncher(t *testing.T) {
	t.Run("data synch success", func(t *testing.T) {
		// given
		ygoClient := &ygo.YgoClientWithCache{}
		ygoCtx := config.YgoContext{}

		// when
		dataSyncher := setup.NewYgoDataSyncher(ygoClient, &ygoCtx)

		// then
		require.NotNil(t, dataSyncher)
	})
}

func TestYgoDataSyncher_Sync(t *testing.T) {
	t.Run("data synch success", func(t *testing.T) {
		// given
		cards := []*model.Card{
			{ID: 111, Name: "Card 111"},
			{ID: 222, Name: "Card 222"},
			{ID: 333, Name: "Card 333"},
		}

		cacheSaverMock := &mocks.CardSaver{}
		cacheSaverMock.On("SaveAllCards", &cards).Return(nil)
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(&cards, nil)

		ygoClient := &ygo.YgoClientWithCache{
			CacheSaver:     cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := setup.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)
	})

	t.Run("fail to save cards", func(t *testing.T) {
		// given
		cards := []*model.Card{
			{ID: 111, Name: "Card 111"},
			{ID: 222, Name: "Card 222"},
			{ID: 333, Name: "Card 333"},
		}

		cacheSaverMock := &mocks.CardSaver{}
		cacheSaverMock.On("SaveAllCards", &cards).Return(assert.AnError)
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(&cards, nil)

		ygoClient := &ygo.YgoClientWithCache{
			CacheSaver:     cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := setup.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)
	})

	t.Run("fail to get all cards", func(t *testing.T) {
		// given
		cacheSaverMock := &mocks.CardSaver{}
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(nil, assert.AnError)

		ygoClient := &ygo.YgoClientWithCache{
			CacheSaver:     cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := setup.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)
	})
}
