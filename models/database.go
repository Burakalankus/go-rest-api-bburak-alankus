package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// PostgreSQL bağlantı bilgilerini oku
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Veritabanına bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully!")
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&Author{}) // Önce Authors oluşturulmalı
	if err != nil {
		log.Fatalf("Migration failed (authors): %v", err)
	}

	err = DB.AutoMigrate(&Book{}) // Sonra Books
	if err != nil {
		log.Fatalf("Migration failed (books): %v", err)
	}

	err = DB.AutoMigrate(&Review{}) // En son Reviews
	if err != nil {
		log.Fatalf("Migration failed (reviews): %v", err)
	}

	log.Println("Database migrated successfully!")
}
