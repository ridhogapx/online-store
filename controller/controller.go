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

func Setup(app *fiber.App, q *db.Queries, secret []byte) *Controller {
	return &Controller{
		App:    app,
		Q:      q,
		Secret: secret,
	}
}

func (controller *Controller) Routes() {
	v1 := controller.App.Group("/api/v1")

	// Customer Route
	v1.Post("/register", controller.Register)
	v1.Post("/login", controller.Login)

	// Product Route
	v1.Get("/products", controller.GetProductByCategory)

	// Shopping cart route
	v1.Post("/carts", controller.CreateCart)
	v1.Get("/carts", controller.FindCarts)
	v1.Delete("/carts/:cart_id", controller.DeleteCart)
}
