package config

import (
	"encoding/json"
	"os"
)

// Config represents the configuration of the application
type Config struct {
	DatabaseURL string `json:"database_url"`
	RabbitMQURL string `json:"rabbitmq_url"`
	ServerURL   string `json:"server_url"`
}

// LoadConfig loads the configuration from a file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
