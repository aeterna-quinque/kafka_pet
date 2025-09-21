package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Greet(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("Welcome!")
}
