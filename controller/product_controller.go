package controller

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func (controller *Controller) GetProductByCategory(c *fiber.Ctx) {
	// Get query category
	category := c.Query("category")

	categoryNum, err := strconv.Atoi(category)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(&Response{
			Message: "Query string for category must be number",
			Status:  "fail",
		})
		return
	}

	res, err := controller.Q.FindProductByCategory(context.Background(), int64(categoryNum))

	if err != nil {
		c.Status(http.StatusNotFound)
		c.JSON(&Response{
			Message: "Data is not found",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusOK)
	c.JSON(res)

}

func (controller *Controller) CreateCategory(c *fiber.Ctx) {
	var bodyRequest CreateCategoryRequest

	c.BodyParser(&bodyRequest)

	res, err := controller.Q.CreateCategory(context.Background(), bodyRequest.CategoryName)

	if err != nil {
		c.Status(http.StatusInternalServerError)

		c.JSON(Response{
			Message: "Failed to create category",
			Status:  "fail",
		})

		return
	}

	c.Status(http.StatusCreated)
	c.JSON(res)
}

func (controller *Controller) CreateProduct(c *fiber.Ctx) {
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
		c.JSON(Response{
			Message: "Failed to create product",
			Status:  "fail",
		})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(res)
}
