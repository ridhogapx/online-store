package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (controller *Controller) CreateCart(c *fiber.Ctx) error {
	var bodyRequest CreateCartRequest

	c.BodyParser(&bodyRequest)

	authorization := c.Get("Authorization")

	res, err := utils.DecodeToken(authorization, controller.Secret)

	// Check if token JWT is valid or not
	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(Response{
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
		return c.JSON(Response{
			Message: "Failed to create shopping cart",
			Status:  "fail",
		})

	}

	c.Status(http.StatusCreated)
	return c.JSON(Response{
		Message: "Product is added into cart!",
		Status:  "success",
	})

}

func (controller *Controller) FindCarts(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	// Validate authorization
	res, err := utils.DecodeToken(authorization, controller.Secret)

	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(Response{
			Message: "Forbidden",
			Status:  "fail",
		})

	}

	// Find Cart by Customer Name
	cart, err := controller.Q.FindCart(context.Background(), res.CustomerID)

	if err != nil {
		c.Status(http.StatusOK)
		return c.JSON(Response{
			Message: "Customer is not yet adding product into shopping cart",
			Status:  "success",
		})

	}

	quantity := len(cart)

	c.Status(http.StatusOK)
	return c.JSON(CartResponse{
		TotalQuantity: int64(quantity),
		Products:      cart,
	})

}

func (controller *Controller) DeleteCart(c *fiber.Ctx) error {
	param_id := c.Params("cart_id")
	authorization := c.Get("authorization")

	// Validate authorization
	_, err := utils.DecodeToken(authorization, controller.Secret)

	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(Response{
			Message: "Forbidden",
			Status:  "fail",
		})

	}

	// If authorization is valid, do DELETE operation
	err = controller.Q.DeleteCart(context.Background(), param_id)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(Response{
			Message: "Failed to delete cart items",
			Status:  "fail",
		})

	}

	c.Status(http.StatusOK)
	return c.JSON(Response{
		Message: "Cart successfully deleted!",
		Status:  "success",
	})

}
