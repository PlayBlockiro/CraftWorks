package server

import (
	"fmt"
	"log"
	"net"

	"github.com/PlayBlockiro/CraftWorks/config"
	"github.com/PlayBlockiro/CraftWorks/utils"
)

func Start(cfg *config.Config) {
	address := fmt.Sprintf(":%d", cfg.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	utils.LogInfo("Server started on " + address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.LogError("Connection error: " + err.Error())
			continue
		}

		go handleConnection(conn, cfg)
	}
}

func handleConnection(conn net.Conn, cfg *config.Config) {
	defer conn.Close()

	token, err := Authenticate(conn)
	if err != nil {
		utils.LogError("Authentication failed: " + err.Error())
		return
	}

	utils.LogInfo("User authenticated with token: " + token)

	if cfg.UseCloudflare {
		if err := HandleCloudflare(conn, cfg); err != nil {
			utils.LogError("Cloudflare verification failed: " + err.Error())
			return
		}
	}

	utils.LogInfo("Client connection established.")
}
