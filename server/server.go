package main

import (
	"github.com/gofiber/fiber/v2"
)

type SomeStruct struct {
	Name string
	Age  uint8
}

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	api.Get("/user", func(c *fiber.Ctx) error {
		data := SomeStruct{
			Name: "Super Admin",
			Age:  24,
		}

		return c.Status(200).JSON(data)
	})

	app.Listen(":8080")
}
