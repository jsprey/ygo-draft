package config_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/config"
	"ygodraft/backend/model/mocks"
)

func TestReadConfig(t *testing.T) {
	t.Run("read config successfully", func(t *testing.T) {
		// when
		c, err := config.ReadConfig("testdata/config.yaml")

		// then
		assert.NoError(t, err)
		assert.Equal(t, 8080, c.Port)
		assert.Equal(t, "DEBUG", c.LogLevel)
		assert.Equal(t, false, c.SyncAtStartup)
		assert.Equal(t, "/", c.ContextPath)
		assert.Equal(t, "http://localhost:5432", c.DatabaseContext.DatabaseUrl)
		assert.Equal(t, "admin", c.DatabaseContext.Username)
		assert.Equal(t, "admin", c.DatabaseContext.Password)
	})

	t.Run("error on reading config with a path to a non existing file", func(t *testing.T) {
		// when
		_, err := config.ReadConfig("testdata/doesnotexist")

		// then
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not find configuration")
	})

	t.Run("error on reading a config file containing invalid yaml", func(t *testing.T) {
		// when
		_, err := config.ReadConfig("testdata/config-error.yml")

		// then
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshal errors")
	})
}

func Test_NewYgoContext(t *testing.T) {
	t.Run("create new context", func(t *testing.T) {
		// given
		ygoClientMock := &mocks.YgoClient{}

		// when
		ctx, err := config.NewYgoContext("testdata/config.yaml", ygoClientMock)

		// then
		require.NoError(t, err)
		assert.NotNil(t, ctx)
	})

	t.Run("new context generation fail because of error when reading file", func(t *testing.T) {
		// given
		ygoClientMock := &mocks.YgoClient{}

		// when
		ctx, err := config.NewYgoContext("non_existent_file", ygoClientMock)

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to read the configuration")
		assert.Nil(t, ctx)
	})
}
