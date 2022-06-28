package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gowlfer/gowlfer/internal/routes"
	"github.com/gowlfer/gowlfer/internal/utils/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("Loading dotenv file...")
	godotErr := godotenv.Load()

	if godotErr != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	log.Println("Setting up CORS...")
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:3000/",
			AllowCredentials: true,
			AllowHeaders:     "Origin, Content-Type, Accept",
		}))

	log.Println("Setting up routes")
	routeError := routes.SetupRoutes(app)

	if routeError != nil {
		log.Fatalf("Error setting up routes: %v", routeError)
	}

	log.Println("Starting up database...")
	dbError := database.StartDB()

	if dbError != nil {
		log.Fatalf("failed to connect to database: %v", dbError)
	}

	// Create table if it doesn't exist
	log.Println("Creating Default Tables if they don't exist")
	//Users.CreateTable()

	app.Listen(":8080")

}
