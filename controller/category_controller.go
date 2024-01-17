package controller

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber"
)

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
