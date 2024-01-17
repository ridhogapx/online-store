package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func (controller *Controller) CreateTransaction(c *fiber.Ctx) {
	authorization := c.Get("Authorization")

	// Check if authorization is valid
	res, err := utils.DecodeToken(authorization, controller.Secret)

	if err != nil {
		c.Status(http.StatusForbidden)
		c.JSON(Response{
			Message: "Forbidden",
			Status:  "fail",
		})

		return
	}

	// We need to get list of carts that user have been added
	cart, err := controller.Q.FindCart(context.Background(), res.CustomerID)

	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(Response{
			Message: "Failed to create transaction because customer is not having cart items",
			Status:  "fail",
		})

		return
	}

	var totalPrice int64

	// Insert into database retrieved items in cart
	for _, items := range cart {
		totalPrice += items.Price
	}

	total, err := controller.Q.CreateTransactionReport(context.Background(), db.CreateTransactionReportParams{
		TransactionID: uuid.NewString(),
		CustomerID:    res.CustomerID,
		TotalPrice:    totalPrice,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(Response{
			Message: "Failed to create transaction",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(CreateTransactionResponse{
		Products:      cart,
		TransactionID: total.TransactionID,
		TotalPrice:    totalPrice,
		CreatedAt:     total.CreatedAt,
	})
}
