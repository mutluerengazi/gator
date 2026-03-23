package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// configFileName is the name of the config file in the home directory.
const configFileName = ".gatorconfig.json"

// Config represents the structure of the JSON config file.
// The struct tags (e.g. `json:"db_url"`) tell the JSON encoder/decoder
// which JSON key maps to which Go field.
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the config file from the home directory and returns a Config.
func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// SetUser sets the current_user_name field on the config and writes it to disk.
// The receiver is a pointer (*Config) so this method can modify the struct.
func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(*cfg)
}

// getConfigFilePath returns the full path to the config file.
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// write serializes the Config to JSON and writes it to the config file.
func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty-print the JSON
	return encoder.Encode(cfg)
}
