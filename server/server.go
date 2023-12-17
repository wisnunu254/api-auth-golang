package server

import (
	"github.com/gofiber/fiber/v2"
)

func Server() {

	// Start the fiber server
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrBadGateway.Code).JSON(fiber.Map{
			"message": "Error",
			"error":   "Bad Gateway",
		})

	})
	app.Listen(":3000")
}
