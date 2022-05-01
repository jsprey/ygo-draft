package config_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/config"
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
		assert.Equal(t, "ygodraft", c.DatabaseContext.DatabaseName)
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
		// when
		ctx, err := config.NewYgoContext("testdata/config.yaml")

		// then
		require.NoError(t, err)
		assert.NotNil(t, ctx)
	})

	t.Run("new context generation fail because of error when reading file", func(t *testing.T) {
		// when
		ctx, err := config.NewYgoContext("non_existent_file")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to read the configuration")
		assert.Nil(t, ctx)
	})
}

func TestDbContext_GetConnectionUrl(t *testing.T) {
	t.Run("correct connection format", func(t *testing.T) {
		// given
		dbConnection := config.DbContext{
			DatabaseUrl:  "localhost:1234",
			DatabaseName: "myDB",
			Username:     "user",
			Password:     "pass",
		}

		// when
		connectionURL := dbConnection.GetConnectionUrl()

		// then
		assert.Equal(t, "postgres://user:pass@localhost:1234/myDB", connectionURL)
	})
}
