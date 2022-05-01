package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"ygodraft/backend/model"
)

// YgoContext contains various configuration values used while running ygo draft.
type YgoContext struct {
	Port            int             `yaml:"port"`
	LogLevel        string          `yaml:"log_level"`
	ContextPath     string          `yaml:"context_path"`
	SyncAtStartup   bool            `yaml:"sync_at_startup"`
	DataClient      model.YgoClient `yaml:"data_store"`
	DatabaseContext DbContext       `yaml:"database_context"`
}

// DbContext contains information about the database to use.
type DbContext struct {
	DatabaseUrl string `yaml:"database_url"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
}

// NewYgoContext creates a new ygo context.
func NewYgoContext(configPath string, client model.YgoClient) (*YgoContext, error) {
	context, err := ReadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the configuration at [%s]: %w", configPath, err)
	}

	context.DataClient = client
	return context, nil
}

// ReadConfig read all provided values from the given configuration file.
func ReadConfig(path string) (*YgoContext, error) {
	config := &YgoContext{}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find configuration at %s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration %s: %w", path, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration %s: %w", path, err)
	}

	return config, nil
}
