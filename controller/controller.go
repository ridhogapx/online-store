package controller

import (
	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber"
)

type Controller struct {
	App    *fiber.App
	Q      *db.Queries
	Secret []byte
}

func Setup(app *fiber.App, q *db.Queries) *Controller {
	return &Controller{
		App: app,
		Q:   q,
	}
}

func (controller *Controller) Routes() {
	v1 := controller.App.Group("/api/v1")

	// Customer Route
	v1.Post("/register", controller.Register)

	// Product Route
	v1.Get("/product", controller.GetProductByCategory)
}
