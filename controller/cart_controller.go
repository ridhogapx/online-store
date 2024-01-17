package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func (controller *Controller) CreateCart(c *fiber.Ctx) {
	var bodyRequest CreateCartRequest

	c.BodyParser(&bodyRequest)

	authorization := c.Get("Authorization")

	res, err := utils.DecodeToken(authorization, controller.Secret)

	// Check if token JWT is valid or not
	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(Response{
			Message: "Authorization is not valid",
			Status:  "fail",
		})
	}

	// If token is valid, insert into database
	_, err = controller.Q.CreateCart(context.Background(), db.CreateCartParams{
		CartID:     uuid.NewString(),
		CustomerID: res.CustomerID,
		ProductID:  bodyRequest.ProductID,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(Response{
			Message: "Failed to create shopping cart",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(Response{
		Message: "Product is added into cart!",
		Status:  "success",
	})
}

func (controller *Controller) FindCarts(c *fiber.Ctx) {
	authorization := c.Get("Authorization")

	// Validate authorization
	res, err := utils.DecodeToken(authorization, controller.Secret)

	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(Response{
			Message: "Forbidden",
			Status:  "fail",
		})
		return
	}

	// Find Cart by Customer Name
	cart, err := controller.Q.FindCart(context.Background(), res.CustomerID)

	if err != nil {
		c.Status(http.StatusOK)
		c.JSON(Response{
			Message: "Customer is not yet adding product into shopping cart",
			Status:  "success",
		})
		return
	}

	quantity := len(cart)
	var total_price int64

	for _, items := range cart {
		total_price += items.Price
	}

	c.Status(http.StatusOK)
	c.JSON(CartResponse{
		TotalQuantity: int64(quantity),
		TotalPrice:    total_price,
		Products:      cart,
	})
}
