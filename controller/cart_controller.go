package controller

import "github.com/gofiber/fiber"

func (controller *Controller) CreateCart(c *fiber.Ctx) {
	var bodyRequest CreateCartRequest

	c.BodyParser(&bodyRequest)

	//authorization := c.Get("Authorization")

}
