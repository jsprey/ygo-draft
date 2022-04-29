package ygoprodeck_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"ygodraft/backend/ygoprodeck"
	"ygodraft/backend/ygoprodeck/mocks"
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
		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := &mocks.HttpClient{}
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
		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(assert.AnError)
		httpClient := &mocks.HttpClient{}
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
		limiter := &mocks.Limiter{}
		limiter.Mock.On("Wait", mock.Anything).Return(nil)
		httpClient := &mocks.HttpClient{}
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
