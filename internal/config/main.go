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

var CONFIG_FILE_NAME string = ".gatorconfig.json"

func Read() (*Config, error) {
	cwd, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to find user home directory: %v", err)
	}

	targetFile := path.Join(cwd, CONFIG_FILE_NAME)
	fd, err := os.Open(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", targetFile, err)
	}

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
