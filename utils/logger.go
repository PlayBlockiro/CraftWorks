package utils

import (
	"fmt"
	"time"
)

func LogInfo(message string) {
	fmt.Printf("[%s] INFO: %s\n", time.Now().Format(time.RFC3339), message)
}

func LogError(message string) {
	fmt.Printf("[%s] ERROR: %s\n", time.Now().Format(time.RFC3339), message)
}
