package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

func StartDB() error {

	databaseConfig := DBConfig{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}

	dbSr := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable", databaseConfig.DBUser, databaseConfig.DBPass, databaseConfig.DBPort, databaseConfig.DBName)

	open, err := gorm.Open(postgres.Open(dbSr), &gorm.Config{})

	DB = open

	return err
}
