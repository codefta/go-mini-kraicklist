package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	return os.Getenv(key)
}