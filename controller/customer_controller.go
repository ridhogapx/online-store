package controller

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
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
		fmt.Println("Error:", err)
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

func (controller *Controller) Login(c *fiber.Ctx) {
	var bodyRequest LoginRequest

	c.BodyParser(&bodyRequest)

	// Check if user is registered.
	res, err := controller.Q.FindCustomerByEmail(context.Background(), bodyRequest.Email)

	if err != nil {
		c.Status(http.StatusNotFound)
		c.JSON(Response{
			Message: "Customer is not registered",
			Status:  "fail",
		})
		return
	}

	// Compare password in hash
	pass := []byte(bodyRequest.Password)
	hashed := []byte(res.Password)

	err = bcrypt.CompareHashAndPassword(hashed, pass)

	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(Response{
			Message: "Password is incorrect",
			Status:  "fal",
		})
		return
	}

	// Generate token if user is success login
	token, err := utils.GenerateToken(utils.Payload{
		CustomerID:   res.CustomerID,
		CustomerName: res.CustomerName,
		Email:        res.Email,
		Secret:       controller.Secret,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(Response{
			Message: "Failed to generate token authorization",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusOK)
	c.JSON(LoginResponse{
		Token: token,
	})
}
