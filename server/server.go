package main

import (
	"github.com/after-friday/vx-blog/controllers/authentication"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	auth := api.Group("/authentication")
	auth.Post("/login", authentication.Login)
	auth.Post("/register", authentication.Register)

	app.Listen(":8080")
}
