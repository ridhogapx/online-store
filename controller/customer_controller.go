package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func (controller *Controller) Register(c *fiber.Ctx) {
	var bodyRequest RegisterRequest

	c.BodyParser(&bodyRequest)

	_, err := controller.Q.CreateCustomer(context.Background(), db.CreateCustomerParams{
		CustomerID:      uuid.New().String(),
		CustomerName:    bodyRequest.Name,
		CustomerAddress: bodyRequest.Address,
		Email:           bodyRequest.Email,
		Password:        bodyRequest.Password,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(Response{
			Message: "Failed to register customer because internal server error",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(Response{
		Message: "Successfully register user!",
		Status:  "success",
	})
}
