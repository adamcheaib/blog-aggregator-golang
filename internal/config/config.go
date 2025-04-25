package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	jsonPath, err := getJsonFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		return Config{}, err
	}

	currentConfig := Config{}
	if err := json.Unmarshal(jsonFile, &currentConfig); err != nil {
		return Config{}, err
	}

	return currentConfig, nil
}

func (c *Config) SetUser(username string) error {
	jsonPath, err := getJsonFilePath()
	if err != nil {
		return err
	}

	c.Current_user_name = username
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}

	if err = os.WriteFile(jsonPath, data, 0600); err != nil {
		return err
	}

	return nil
}

func getJsonFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFileName := ".gatorconfig.json"

	path += fmt.Sprintf("/Documents/GitHub/blog-aggregator-golang/%v", configFileName)
	return path, nil
}
