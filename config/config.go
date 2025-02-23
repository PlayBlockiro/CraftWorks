package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port             int    `json:"port"`
	UseCloudflare    bool   `json:"use_cloudflare"`
	CloudflareAPIKey string `json:"cloudflare_api_key,omitempty"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
