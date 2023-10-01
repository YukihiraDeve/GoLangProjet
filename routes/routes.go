package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Définir une route pour le chemin "/api"
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Im a GET request!")
	})

	// Groupe de routes pour le chemin "/test"
	testGroup := app.Group("/test")
	testGroup.Get("/foo", handler1)
	testGroup.Get("/bar", handler2)
}

func handler1(c *fiber.Ctx) error {
	return c.SendString("Handling GET request on /test/foo!")
}

func handler2(c *fiber.Ctx) error {
	return c.SendString("Handling GET request on /test/bar!")
}
