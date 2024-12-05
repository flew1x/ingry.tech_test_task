package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var cfg *koanf.Koanf

type Config struct {
	RestConfig     IRESTConfig
	PostgresConfig IPostgresConfig
}

func NewConfig() *Config {
	return &Config{}
}

// InitConfig loads the configuration from a YAML file.
//
// Parameters:
// - configPath: the directory where the configuration file is stored.
// - configFile: the name of the configuration file.
//
// Returns:
// - error: an error if the configuration can't be loaded.
func (c *Config) InitConfig(configPath, configFile string) error {
	cfg = koanf.New(".")

	filePath := filepath.Join(configPath, configFile)

	config := file.Provider(filePath)

	if err := cfg.Load(config, yaml.Parser()); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	c.RestConfig = NewRestConfig()
	c.PostgresConfig = NewPostgresConfig()

	return nil
}

// MustStringFromEnv returns the value of the environment variable or panics if the environment variable is not set.
//
// Parameters:
// - field: the name of the environment variable to retrieve.
//
// Returns:
// - string: the value of the environment variable.
func mustStringFromEnv(field string) string {
	envValue := os.Getenv(field)

	if envValue == "" {
		panic(fmt.Sprintf("environment variable %s is not set", field))
	}

	return envValue
}

// MustUnmarshal unmarshals the field in the config or panics if the field does not exist.
//
// Parameters:
// - field: the name of the field to retrieve.
// - v: the pointer to the struct to unmarshal into.
//
// Returns:
// - nil
func mustUnmarshalStruct(field string, v any) {
	if err := cfg.Unmarshal(field, v); err != nil {
		panic(err)
	}
}
