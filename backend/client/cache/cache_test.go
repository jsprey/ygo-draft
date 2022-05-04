package cache_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/client/cache/mocks"
	"ygodraft/backend/query"
)

func TestNewYgoCache(t *testing.T) {
	t.Run("create new cache", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}

		// when
		myCache, err := cache.NewYgoCache(dbMock)

		// then
		require.NoError(t, err)
		assert.NotNil(t, myCache)
	})

	t.Run("fail the creation because of an invalid template", func(t *testing.T) {
		// given
		originalTemplate := query.TemplateContentSelectCardByID
		defer func() { query.TemplateContentSelectCardByID = originalTemplate }()
		query.TemplateContentSelectCardByID = "SELECT * FROM public.cards {{}{}}"

		dbMock := &mocks.DatabaseClient{}

		// when
		_, err := cache.NewYgoCache(dbMock)

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create new sql query templater")
	})
}
