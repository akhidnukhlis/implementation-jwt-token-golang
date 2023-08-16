package helpers

import (
	"log"
	"os"
)

// GetCurrentDirectory read current directory
func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to read current directory: %s", err)
	}

	return dir
}
