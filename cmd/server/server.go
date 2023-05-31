package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	registerRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
