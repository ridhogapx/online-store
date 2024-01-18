package controller

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) CreateCategory(c *fiber.Ctx) error {
	var bodyRequest CreateCategoryRequest

	c.BodyParser(&bodyRequest)

	res, err := controller.Q.CreateCategory(context.Background(), bodyRequest.CategoryName)

	if err != nil {
		c.Status(http.StatusInternalServerError)

		return c.JSON(Response{
			Message: "Failed to create category",
			Status:  "fail",
		})

	}

	c.Status(http.StatusCreated)
	return c.JSON(res)

}

func (controller *Controller) FindCategories(c *fiber.Ctx) error {
	res, err := controller.Q.FindAllCategories(context.Background())

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(Response{
			Message: "Category is not empty",
			Status:  "fail",
		})
	}

	c.Status(http.StatusOK)
	return c.JSON(res)

}
