package routes

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"timestamp": time.Now(),
	})
}

func Controller(app *fiber.App) {
	app.Get("/", GetIndex)
}
