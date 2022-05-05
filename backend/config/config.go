package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	StageProductive  = "production"
	StageDevelopment = "development"
)

// YgoContext contains various configuration values used while running ygo draft.
type YgoContext struct {
	Port            int       `yaml:"port"`
	LogLevel        string    `yaml:"log_level"`
	ContextPath     string    `yaml:"context_path"`
	SyncAtStartup   bool      `yaml:"sync_at_startup"`
	Stage           string    `yaml:"stage"`
	DatabaseContext DbContext `yaml:"database_context"`
}

// DbContext contains information about the database to use.
type DbContext struct {
	DatabaseUrl  string `yaml:"database_url"`
	DatabaseName string `yaml:"database_name"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

func (dc *DbContext) GetConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", dc.Username, dc.Password, dc.DatabaseUrl, dc.DatabaseName)
}

// NewYgoContext creates a new ygo context.
func NewYgoContext(configPath string) (*YgoContext, error) {
	context, err := ReadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the configuration at [%s]: %w", configPath, err)
	}

	if context.Stage == StageProductive {
		//gin.SetMode(gin.ReleaseMode)
	}

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
