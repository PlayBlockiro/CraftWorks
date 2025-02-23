package server

import (
	"errors"
	"net"
	"time"

	"github.com/PlayBlockiro/CraftWorks/config"
)

func HandleCloudflare(conn net.Conn, cfg *config.Config) error {
	if cfg.UseCloudflare {
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			return errors.New("cloudflare handshake failed")
		}
	}

	return nil
}
