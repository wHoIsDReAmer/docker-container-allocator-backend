package main

import (
	"dca-backend/routes"
	"dca-backend/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	util.LoadEnv(".env")

	app := fiber.New(fiber.Config{
		//DisableStartupMessage: true,
	})

	routes.Controller(app)

	app.Listen(":90")
}
