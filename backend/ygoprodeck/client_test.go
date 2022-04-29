package ygoprodeck_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/ygoprodeck"
	"ygodraft/backend/ygoprodeck/mocks"
)

func TestNewYgoProDeckClient(t *testing.T) {
	t.Run("create new client", func(t *testing.T) {
		// when
		client := ygoprodeck.NewYgoProDeckClient()

		// then
		assert.NotNil(t, client)
		assert.NotNil(t, client.Client)
	})
}

func TestYgoProDeckClient_GetAllCards(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		// given
		myCards := struct {
			Data []model.Card `json:"data,omitempty"`
		}{
			Data: []model.Card{
				{Name: "mycard1"},
				{Name: "mycard2"},
				{Name: "mycard3"},
			},
		}
		myCardsBytes, err := json.Marshal(myCards)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/cardinfo.php" {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(myCardsBytes)
				require.NoError(t, err)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := http.DefaultClient
		client := &ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}
		ygoClient := ygoprodeck.YgoProDeckClient{BaseUrl: server.URL, Client: client}

		// when
		cards, err := ygoClient.GetAllCards()

		// then
		require.NoError(t, err)
		assert.Len(t, *cards, 3)
		assert.Equal(t, "mycard1", (*cards)[0].Name)
		assert.Equal(t, "mycard2", (*cards)[1].Name)
		assert.Equal(t, "mycard3", (*cards)[2].Name)
	})

	t.Run("errors on get request", func(t *testing.T) {
		// given
		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := &mocks.HttpClient{}
		httpClient.Mock.On("Get", mock.Anything).Return(nil, assert.AnError)
		client := &ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}
		ygoClient := ygoprodeck.YgoProDeckClient{Client: client}

		// when
		cards, err := ygoClient.GetAllCards()

		// then
		require.Error(t, err)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, cards)
	})

	t.Run("errors on unmarshal", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/cardinfo.php" {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte("not expected json data"))
				require.NoError(t, err)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := http.DefaultClient
		client := &ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}
		ygoClient := ygoprodeck.YgoProDeckClient{BaseUrl: server.URL, Client: client}

		// when
		cards, err := ygoClient.GetAllCards()

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal body")
		assert.Nil(t, cards)
	})
}
