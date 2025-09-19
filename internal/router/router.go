package router

import (
	"kafka-pet/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(controller *controller.Controller) *fiber.App {
	r := fiber.New()

	r.All("/", controller.Greet)

	return r
}
