package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (controller *Controller) Register(c *fiber.Ctx) {
	var bodyRequest RegisterRequest

	c.BodyParser(&bodyRequest)

	// Check existing customer
	// If customer is already exist, then return fail response
	_, err := controller.Q.FindCustomerByEmail(context.Background(), bodyRequest.Email)

	if err == nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(Response{
			Message: "Customer is already registered",
			Status:  "fail",
		})

		return
	}

	pass := []byte(bodyRequest.Password)
	hash, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	_, err = controller.Q.CreateCustomer(context.Background(), db.CreateCustomerParams{
		CustomerID:      uuid.New().String(),
		CustomerName:    bodyRequest.Name,
		CustomerAddress: bodyRequest.Address,
		Email:           bodyRequest.Email,
		Password:        string(hash),
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
