package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Routers (Handlers)

	app.Post("/api/users", func(c *fiber.Ctx) error {
		return c.SendString("POST /api/users")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return c.SendString("GET /api/users")
	})

	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("GET /api/users/%s", c.Params("id")))
	})

	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("GET /api/users/%s", c.Params("id")))
	})

	app.Put("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("PUT /api/users/%s", c.Params("id")))
	})

	log.Fatal(app.Listen(":9234"))
}
