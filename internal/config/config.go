package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
}

func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("geting home path: %w", err)
	}
	return filepath.Join(home, ".particard"), nil
}

func GetConfigPath() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func SaveConfig(config Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	var mkdirPerm os.FileMode = 0o755

	if err := os.MkdirAll(filepath.Dir(configPath), mkdirPerm); err != nil {
		return fmt.Errorf("save config: can not create a directories '%s'with perm '%s': %w", filepath.Dir(configPath), mkdirPerm.String(), err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("save config: marshal data in json: %W", err)
	}

	var writePerm os.FileMode = 0o644

	if err := os.WriteFile(configPath, data, 0o644); err != nil {
		return fmt.Errorf("save config: write data in file '%s' with perm '%s': %w", configPath, writePerm.String(), err)
	}

	return nil
}

func LoadConfig() (Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return Config{}, fmt.Errorf("load config: get config path: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("load config: read data from file '%s': %w", configPath, err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("load config: unmarshal data from json: %w", err)
	}

	return config, nil
}
