package main

import (
	"github.com/gowlfer/gowlfer/app/models"
	"github.com/gowlfer/gowlfer/internal/utils/database"
	"github.com/joho/godotenv"
	"log"
)

// Our tables array, just add the models you want to create here
var tables = []any{
	&models.GowlferUser{},
	&models.GowlferCompany{},
	&models.GowlferCompanyType{},
}

func main() {
	log.Println("Loading dotenv file...")
	godotErr := godotenv.Load()

	if godotErr != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Starting DB...")
	dbError := database.StartDB()
	if dbError != nil {
		log.Fatalf("failed to connect to database: %v", dbError)
	}
	log.Println("Creating Tables...")

	// We will loop through our tables array and create the tables
	for table := range tables {
		database.DB.AutoMigrate(tables[table])
	}
}
