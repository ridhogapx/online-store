package controller

import (
	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber"
)

type Controller struct {
	App *fiber.App
	Q   *db.Queries
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Setup(app *fiber.App, q *db.Queries) *Controller {
	return &Controller{
		App: app,
		Q:   q,
	}
}

func (controller *Controller) Routes() {
	v1 := controller.App.Group("/api/v1")
	v1.Get("/product", controller.GetProductByCategory)
}
