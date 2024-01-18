package main

import (
	"fmt"
	"os"

	"github.com/RageNeko26/online-store/controller"
	"github.com/RageNeko26/online-store/db"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// Load up environment variable
	SECRET := os.Getenv("SECRET")

	// Initialize configuration
	q := db.NewDBConnection(&db.DBCredentials{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	})

	controller := controller.Setup(app, q, []byte(SECRET))
	// Initialize route
	controller.Routes()

	// Internal logging
	fmt.Println("Server is running")

	app.Listen(":3000")
}
