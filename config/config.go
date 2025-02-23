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
	
	// We will try to save (create) the config file each time server.go is ran.
	
	// Perhaps I should rename SaveConfig to CreateConfig instead and make a new function to actually save???
	SaveConfig(filePath);
	
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

// If a config doesnt exist, create it, and then subsequently save a config file
func SaveConfig(filePath string) (*Config, error) {
	// variables for our file, and make sure to check for errors when creating our file.
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	
	// defer closes file at the end of our function, very nice.
	defer file.Close()
	
	var cfg Config
	
	// Create our config with our default data
	config := Config {
		Port: 8000,
		UseCloudflare: false,
		CloudflareAPIKey: "key",
	}
	
	/*
	 * Convert our config data to json
	 * Why is it called Marshal? I have no clue.
	 * 
	 * https://pkg.go.dev/encoding/json#MarshalIndent
	 * 
	 */
	
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return nil, err
	}
	
	/*
	 * Write out a string into our new shiny file.
	 * 
	 * https://pkg.go.dev/io/fs#FileMode
	 * 
	 */
	err = os.WriteFile(filePath, bytes, 0644)
	if err != nil {
		return nil, err
	}
	
	// return our delicious config file!!!^!6!!^6!^!^^^!116616!6
	return &cfg, nil
	
}
