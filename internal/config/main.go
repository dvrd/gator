package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	cwd, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to find user home directory: %v", err)
	}

	targetFile := path.Join(cwd, configFileName)

	return targetFile, nil
}

func Read() (*Config, error) {
	targetFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	fd, err := os.Open(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", targetFile, err)
	}
	defer fd.Close()

	data, err := io.ReadAll(fd)
	if err != nil {
		return nil, fmt.Errorf("failed to read file descriptor: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(cfg)
}

func write(cfg *Config) error {
	targetFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	fd, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", targetFile, err)
	}
	defer fd.Close()

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to parse configuration into json: %v", err)
	}

	_, err = fd.Write(jsonData)
	if err != nil {
		return fmt.Errorf("[config.write] %v", err)
	}

	return nil
}
