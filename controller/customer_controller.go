package controller

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Add Customer
//
//	@Summary Add new customer data
//	@Description 	Register new customer
//	@Tags			Customer
//	@Accept			json
//	@Produce 		json
//	@Param 			category body				RegisterRequest		true		"add customer"
//	@Success		201							{object} Response
//	@Failure		500							{object} Response
//	@Failure		400							{object} Response
//	@Router			/api/v1/register [post]
func (controller *Controller) Register(c *fiber.Ctx) error {
	var bodyRequest RegisterRequest

	c.BodyParser(&bodyRequest)

	// Check existing customer
	// If customer is already exist, then return fail response
	_, err := controller.Q.FindCustomerByEmail(context.Background(), bodyRequest.Email)

	if err == nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(Response{
			Message: "Customer is already registered",
			Status:  "fail",
		})

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
		return c.JSON(Response{
			Message: "Failed to register customer because internal server error",
			Status:  "fail",
		})

	}

	c.Status(http.StatusCreated)
	return c.JSON(Response{
		Message: "Successfully register user!",
		Status:  "success",
	})

}

// Check Customer
//
//	@Summary Check and authorize customer.
//	@Description 	Authorize customer and returning token.
//	@Tags			Customer
//	@Accept			json
//	@Produce 		json
//	@Param 			category body				LoginRequest		true		"check customer"
//	@Success		200							{object} Response
//	@Failure		500							{object} Response
//	@Failure		404							{object} Response
//	@Failure		400							{object} Response
//	@Router			/api/v1/login [post]
func (controller *Controller) Login(c *fiber.Ctx) error {
	var bodyRequest LoginRequest

	c.BodyParser(&bodyRequest)

	// Check if user is registered.
	res, err := controller.Q.FindCustomerByEmail(context.Background(), bodyRequest.Email)

	if err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(Response{
			Message: "Customer is not registered",
			Status:  "fail",
		})

	}

	// Compare password in hash
	pass := []byte(bodyRequest.Password)
	hashed := []byte(res.Password)

	err = bcrypt.CompareHashAndPassword(hashed, pass)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(Response{
			Message: "Password is incorrect",
			Status:  "fail",
		})

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
		return c.JSON(Response{
			Message: "Failed to generate token authorization",
			Status:  "fail",
		})
	}

	c.Status(http.StatusOK)
	return c.JSON(LoginResponse{
		Token: token,
	})

}
