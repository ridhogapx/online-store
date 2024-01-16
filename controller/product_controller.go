package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
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
