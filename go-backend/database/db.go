package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global DB Connection
var DB *gorm.DB
func ConnectDB() {
	var err error

	// 1. GET config from .env
	// Format Connection String Supabase (Postgres)
	dsn := os.Getenv("DATABASE_URL")

	// 2. Try to connect to Postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to Database: ", err)
	}

	log.Println("✅ Successfully connected to Supabase Database! 🚀")
	log.Println("⚙️  Creating Tables...")

	// 3. Auto Migrate Tables
	err = DB.AutoMigrate(&User{}, &ScanHistory{}, &Like{})
	if err != nil {
		log.Fatal("❌ Failed to Migrate Table: ", err)
	}

	log.Println("✅ Successfully created Users, Scans, & Likes tables! 🎉")
}