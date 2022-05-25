package synch_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/client/ygo/mocks"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/synch"
)

func TestNewYgoDataSyncher(t *testing.T) {
	t.Run("data synch success", func(t *testing.T) {
		// given
		ygoClient := &ygo.YgoClientWithCache{}
		ygoCtx := config.YgoContext{}

		// when
		dataSyncher, _ := synch.NewYgoDataSyncher(ygoClient, &ygoCtx)

		// then
		require.NotNil(t, dataSyncher)
	})
}

func TestYgoDataSyncher_Sync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/small.png" || r.URL.Path == "/big.png" {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("Test"))
			require.NoError(t, err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))

	t.Run("data synch success", func(t *testing.T) {
		// given
		cards := []*model.Card{
			{ID: 111, Name: "Card 111", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/big.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/small.png", server.URL)}}},
			{ID: 222, Name: "Card 222", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/big.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/small.png", server.URL)}}},
			{ID: 333, Name: "Card 333", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/big.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/small.png", server.URL)}}},
		}

		cacheSaverMock := &mocks.CardSaver{}
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(&cards, nil)
		cacheRetrieverMocker.On("GetCard", 111).Return(cards[0], nil)
		cacheRetrieverMocker.On("GetCard", 222).Return(cards[1], cache.ErrorCardDoesNotExist.WithParam("222"))
		cacheRetrieverMocker.On("GetCard", 333).Return(cards[2], nil)
		cacheSaverMock.On("SaveCard", cards[1]).Return(nil)

		ygoClient := &ygo.YgoClientWithCache{
			CacheCardSaver: cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := synch.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)

		err = os.RemoveAll(synch.ImagesDirectoryName)
		require.NoError(t, err)
	})

	t.Run("fail to save cards", func(t *testing.T) {
		// given
		cards := []*model.Card{
			{ID: 111, Name: "Card 111"},
			{ID: 222, Name: "Card 222"},
			{ID: 333, Name: "Card 333"},
		}

		cacheSaverMock := &mocks.CardSaver{}
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(&cards, nil)
		cacheRetrieverMocker.On("GetCard", 111).Return(cards[0], nil)
		cacheRetrieverMocker.On("GetCard", 222).Return(cards[1], cache.ErrorCardDoesNotExist.WithParam("222"))
		cacheSaverMock.On("SaveCard", cards[1]).Return(assert.AnError)

		ygoClient := &ygo.YgoClientWithCache{
			CacheCardSaver: cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := synch.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)

		err = os.RemoveAll(synch.ImagesDirectoryName)
		require.NoError(t, err)
	})

	t.Run("fail to get all cards", func(t *testing.T) {
		// given
		cacheSaverMock := &mocks.CardSaver{}
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(nil, assert.AnError)

		ygoClient := &ygo.YgoClientWithCache{
			CacheCardSaver: cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := synch.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)

		err = os.RemoveAll(synch.ImagesDirectoryName)
		require.NoError(t, err)
	})

	t.Run("data synch success", func(t *testing.T) {
		// given
		cards := []*model.Card{
			{ID: 111, Name: "Card 111", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/no.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/no.png", server.URL)}}},
			{ID: 222, Name: "Card 222", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/no.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/no.png", server.URL)}}},
			{ID: 333, Name: "Card 333", Images: []model.CardImage{{ImageURL: fmt.Sprintf("%s/no.png", server.URL), ImageURLSmall: fmt.Sprintf("%s/no.png", server.URL)}}},
		}

		cacheSaverMock := &mocks.CardSaver{}
		cacheRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker := &mocks.CardRetriever{}
		webRetrieverMocker.On("GetAllCards").Return(&cards, nil)
		cacheRetrieverMocker.On("GetCard", 111).Return(cards[0], nil)
		cacheRetrieverMocker.On("GetCard", 222).Return(cards[1], cache.ErrorCardDoesNotExist.WithParam("222"))
		cacheRetrieverMocker.On("GetCard", 333).Return(cards[2], nil)
		cacheSaverMock.On("SaveCard", cards[1]).Return(nil)

		ygoClient := &ygo.YgoClientWithCache{
			CacheCardSaver: cacheSaverMock,
			CacheRetriever: cacheRetrieverMocker,
			WebClient:      webRetrieverMocker,
		}
		ygoCtx := &config.YgoContext{Stage: config.StageDevelopment}
		snycher := synch.YgoDataSyncher{
			Client: ygoClient,
			YgoCtx: ygoCtx,
		}

		// when
		err := snycher.Sync()

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, cacheSaverMock, cacheRetrieverMocker, webRetrieverMocker)

		err = os.RemoveAll(synch.ImagesDirectoryName)
		require.NoError(t, err)
	})
}
