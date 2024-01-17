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
