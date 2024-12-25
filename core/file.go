package core

import (
	"log"
	"os"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory")
	}
	return dir
}