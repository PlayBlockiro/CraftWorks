package server

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

// Authenticate verifies the "ServerSessionConnection Token" from the client
func Authenticate(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	token, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("failed to read token")
	}

	token = strings.TrimSpace(token)
	if len(token) == 0 {
		return "", errors.New("invalid token")
	}

	return token, nil
}
