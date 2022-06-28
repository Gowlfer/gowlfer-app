package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowlfer/gowlfer/app/controllers"
)

func SetupRoutes(app *fiber.App) error {

	// User Routes
	app.Post("/api/v1/register", controllers.RegisterUser)
	app.Post("/api/v1/login", controllers.LoginUser)
	app.Get("/api/v1/user", controllers.GetUser)

	// Company Routes
	app.Get("/api/v1/companies", controllers.GetCompanies)
	app.Get("/api/v1/companies", controllers.GetCompany)
	app.Post("/api/v1/createCompany", controllers.CreateCompany)

	return nil
}
