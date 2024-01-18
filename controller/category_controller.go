package controller

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Add Category
//
//	@Summary Add category product data
//	@Description 	Adding category for product
//	@Tags			Product Category
//	@Accept			json
//	@Produce 		json
//	@Param 			category body				CreateCategoryRequest		true		"add category"
//	@Success		201							{object} db.Category
//	@Failure		500							{object} Response
//	@Router			/api/v1/categories [post]
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

// Get All Categoriy
//
//	@Summary Get product categories
//	@Description 	Listing all of categories product.
//	@Tags			Product Category
//	@Produce 		json
//	@Success		200							{object} []db.Category
//	@Failure		404							{object} Response
//	@Router			/api/v1/categories [get]
func (controller *Controller) FindCategories(c *fiber.Ctx) error {
	res, err := controller.Q.FindAllCategories(context.Background())

	if err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(Response{
			Message: "Category is empty",
			Status:  "fail",
		})
	}

	c.Status(http.StatusOK)
	return c.JSON(res)

}
