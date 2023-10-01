package main

import (
	"log"

	"Golang/database"
	"Golang/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Static("/", "./frontend/dist")

	app.Post("/add", routes.AddUser)
	app.Post("/get", routes.GetUser)
	app.Put("/update", routes.GetAllUsers)
	app.Delete("/delete", routes.DeleteUser)
}

func main() {
	database.ConnectDatabse()

	app := fiber.New(fiber.Config{})

	setUpRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	log.Fatal(app.Listen(":3000"))
}
