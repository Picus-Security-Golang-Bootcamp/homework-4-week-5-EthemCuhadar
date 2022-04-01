package config

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

var envFile = "../.env"

// GetConfig takes parameters from environment file and creates dns to access the database.
func GetConfig() string{
	err := godotenv.Load(envFile)
	if err != nil {
	fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("BOOKSTORE_DB_HOST"),
		os.Getenv("BOOKSTORE_DB_PORT"),
		os.Getenv("BOOKSTORE_DB_USERNAME"),
		os.Getenv("BOOKSTORE_DB_NAME"),
		os.Getenv("BOOKSTORE_DB_PASSWORD"),
	)
	return dataSourceName
}
