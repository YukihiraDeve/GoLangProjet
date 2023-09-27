package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	setLog()
	log.Info("Starting server...")

	app := fiber.New(fiber.Config{})
	app.Static("/", "./public")
	app.Get("/api", func(c *fiber.Ctx) error {
		fmt.Println("GET /api from", c.IP())
		return c.SendString("Im a GET request!")
	})

	app.Route("/test", func(api fiber.Router) {
		api.Get("/foo", handler1).Name("foo")
		api.Get("/bar", handler2).Name("bar")
	}, "test.")
	app.Listen(":80")

}

func handler1(c *fiber.Ctx) error {
	// handler for test/foo
	fmt.Println("GET /test/foo from ", c.IP())
	return c.SendString("I'm a GET request /test/foo!")
}

func handler2(c *fiber.Ctx) error {
	// handler for test/bar
	fmt.Println("GET /test/bar from ", c.IP())
	return c.SendString("I'm a GET request /test/bar!")
}

func setLog() {
	fmt.Println("Settings up log file : ", time.Now().UTC().Format("2006-01-02"))
	logFileName := fmt.Sprintf("./logs/%s.log", time.Now().UTC().Format("2006-01-02"))
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}
