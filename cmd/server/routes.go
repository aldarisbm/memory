package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func registerApiV1Routes(app fiber.Router) {
	// POST /api/v1/memory
	// Create a new memory
	app.Post("/memory", func(c *fiber.Ctx) error {
		return c.JSON(&PostMemoryRes{
			ID: uuid.New(),
		})
	})
}
