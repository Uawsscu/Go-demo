package config

import (
	"fmt"
	"go-demo/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("[database] Error loading .env file")
	}

	// ----------------------------Database Connection---------------------------------
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[database] failed to connect database")
	}

	db.AutoMigrate(&models.Anime{}, &models.Character{}, &models.FileInfo{})
	dropColumn(db)

	fmt.Println("[database] Connected to PostgreSQL successfully!")
	DB = db
}

func dropColumn(db *gorm.DB) {
	// db.Migrator().DropColumn(&models.Character{}, "url")
}
