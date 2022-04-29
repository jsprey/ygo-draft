package config_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ygodraft/backend/config"
)

func TestReadConfig(t *testing.T) {
	c, err := config.ReadConfig("testdata/config.yaml")
	assert.NoError(t, err)
	assert.Equal(t, 1234, c.Port)
}

func TestReadConfig_doesNotExist(t *testing.T) {
	_, err := config.ReadConfig("testdata/doesnotexist")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not find configuration")
}

func TestReadConfig_notYaml(t *testing.T) {
	_, err := config.ReadConfig("testdata/config-error.yml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unmarshal errors")
}
