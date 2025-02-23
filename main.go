package main

import (
	"github.com/PlayBlockiro/CraftWorks/config"
	"github.com/PlayBlockiro/CraftWorks/server"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/main.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Start the server
	server.Start(cfg)
}
