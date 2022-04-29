package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("../testdata/config.yaml")
	assert.NoError(t, err)
	assert.Equal(t, 1234, config.Port)
	assert.Equal(t, "testuser", config.ProjectRepository.Username)
}

func TestReadConfig_doesNotExist(t *testing.T) {
	_, err := ReadConfig("../testdata/doesnotexist")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not find configuration")
}

func TestReadConfig_notYaml(t *testing.T) {
	_, err := ReadConfig("../testdata/config-error.yml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unmarshal errors")
}
