package config

import (
	"log"
    // "os"

    "github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }
}