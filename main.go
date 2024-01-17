package main

import (
	"fmt"
	"os"

	"github.com/RageNeko26/online-store/config"
	"github.com/RageNeko26/online-store/controller"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// Load up environment variable
	DB_SOURCE := os.Getenv("DB_SOURCE")
	SECRET := os.Getenv("SECRET")

	// Initialize configuration
	q := config.NewDBConnection(DB_SOURCE)

	controller := controller.Setup(app, q, []byte(SECRET))
	// Initialize route
	controller.Routes()

	// Internal logging
	fmt.Println("Server is running")

	app.Listen(":3000")
}
