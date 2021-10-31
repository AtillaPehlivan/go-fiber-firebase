package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Loading Environment File")
		return err
	}
	return nil
}

func Get(key string) string {
	return os.Getenv(key)
}
