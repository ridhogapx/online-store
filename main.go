package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Listen(":3000")
}
