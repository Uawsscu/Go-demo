package main

import (
	"fmt"
	"main/src/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load values from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read database connection parameters from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Create the DSN string
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate add/update
	db.AutoMigrate(&models.Anime{}, &models.Character{})
}
