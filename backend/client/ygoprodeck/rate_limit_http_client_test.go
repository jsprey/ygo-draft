package ygoprodeck_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"ygodraft/backend/client/ygoprodeck"
	mocks2 "ygodraft/backend/client/ygoprodeck/mocks"
	"ygodraft/backend/model"
)

func TestNewDefaultRateLimitedClient(t *testing.T) {
	// when
	client := ygoprodeck.NewDefaultRateLimitedClient()

	// then
	assert.NotNil(t, client)
	assert.NotNil(t, client.Client)
	assert.NotNil(t, client.RateLimiter)
}

func TestRLHTTPClient_Do(t *testing.T) {
	t.Run("limiter is called", func(t *testing.T) {
		// given
		limiter := &mocks2.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := &mocks2.HttpClient{}
		httpClient.Mock.On("Do", mock.Anything).Return(&http.Response{}, nil)
		client := ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}

		// when
		do, err := client.Do(&http.Request{})

		// then
		require.NoError(t, err)
		assert.NotNil(t, do)
		mock.AssertExpectationsForObjects(t, limiter)
	})

	t.Run("limiter returns error", func(t *testing.T) {
		// given
		limiter := &mocks2.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(assert.AnError)
		httpClient := &mocks2.HttpClient{}
		client := ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}

		// when
		do, err := client.Do(&http.Request{})

		// then
		require.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, do)
		mock.AssertExpectationsForObjects(t, limiter)
	})

	t.Run("limiter returns error", func(t *testing.T) {
		// given
		limiter := &mocks2.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := &mocks2.HttpClient{}
		httpClient.Mock.On("Do", mock.Anything).Return(nil, assert.AnError)
		client := ygoprodeck.RLHTTPClient{
			Client:      httpClient,
			RateLimiter: limiter,
		}

		// when
		do, err := client.Do(&http.Request{})

		// then
		require.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, do)
		mock.AssertExpectationsForObjects(t, limiter)
	})
}

func TestRLHTTPClient_getJsonFromTarget(t *testing.T) {
	t.Run("retrieve requested type correctly", func(t *testing.T) {
		// given
		testCard := model.Card{ID: 123123, Name: "My Test Card"}
		testCardBytes, _ := json.Marshal(testCard)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/test" {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(testCardBytes)
				require.NoError(t, err)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
		}))
		client := ygoprodeck.NewDefaultRateLimitedClient()

		// when
		var myCard model.Card
		err := client.GetJsonFromTarget(fmt.Sprintf("%s/test", testServer.URL), &myCard)

		// then
		require.NoError(t, err)
		assert.Equal(t, 123123, myCard.ID)
		assert.Equal(t, "My Test Card", myCard.Name)
	})

	t.Run("error wrong data -> error on unmarshal", func(t *testing.T) {
		// given
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/test" {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte("Invalid JSON"))
				require.NoError(t, err)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
		}))
		client := ygoprodeck.NewDefaultRateLimitedClient()

		// when
		var myCard model.Card
		err := client.GetJsonFromTarget(fmt.Sprintf("%s/test", testServer.URL), &myCard)

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal body")
	})
}
