package main

import (
	"golangApi/database"
	"golangApi/routes"

	"github.com/gofiber/fiber"
	"github.com/rs/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Options{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")

}
