package router

import (
	"kafka-pet/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(controller *controller.Controller) *fiber.App {
	r := fiber.New()

	r.All("/", controller.Greet)
	r.Post("/create", controller.CreateUser)
	r.Get("/get", controller.GetUser)
	r.Get("/messages", controller.GetMessages)

	return r
}
