package controller

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (controller *Controller) GetProductByCategory(c *fiber.Ctx) error {
	// Get query category
	category := c.Query("category")

	categoryNum, err := strconv.Atoi(category)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Response{
			Message: "Query string for category must be number",
			Status:  "fail",
		})

	}

	res, err := controller.Q.FindProductByCategory(context.Background(), int64(categoryNum))

	if err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(&Response{
			Message: "Data is not found",
			Status:  "fail",
		})

	}

	c.Status(http.StatusOK)
	return c.JSON(res)

}

func (controller *Controller) CreateProduct(c *fiber.Ctx) error {
	var bodyRequest CreateProductRequest

	c.BodyParser(&bodyRequest)

	res, err := controller.Q.CreateProduct(context.Background(), db.CreateProductParams{
		ProductID:   uuid.NewString(),
		CategoryID:  bodyRequest.CategoryID,
		ProductName: bodyRequest.ProductName,
		Price:       bodyRequest.Price,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(Response{
			Message: "Failed to create product",
			Status:  "fail",
		})

	}

	c.Status(http.StatusCreated)
	return c.JSON(res)
}
