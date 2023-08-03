package main

import (
	"database/sql"
	"fmt"

	"github.com/after-friday/vx-blog/controllers/authentication"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	connectionStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", viper.GetString("db.host"), viper.GetInt("db.port"), viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.database"))

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	db.Close()

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

	app.Listen(viper.GetString("app.port"))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
