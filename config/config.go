package config

import (
	"encoding/json"
	"os"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Server `json:"server"`
}

func Parse() (*Config, error) {
	fileByte, readErr := os.ReadFile("./config/config.json")
	if readErr != nil {
		return nil, readErr
	}

	var config Config
	if err := json.Unmarshal(fileByte, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
