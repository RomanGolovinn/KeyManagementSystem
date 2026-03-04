package main

import (
	"log"
	"os"

	"github.com/RomanGolovinn/KeyManagementSystem/internal/storage/postgres"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn, found := os.LookupEnv("DSN")
	if !found {
		log.Fatal("DSN not found")
	}

	psql, err := postgres.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("Failed to create postgres: %v", err)
	}
	defer psql.DB.Close()
}
