package controller

import (
	"kafka-pet/internal/dto"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		c.logger.Error("Couldn't parse request body", zap.ByteString("uri", ctx.Request().RequestURI()), zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.ErrBadRequest, "message": "couldn't parse request"})
	}

	id, err := c.service.CreateUser(c.ctx, &req)
	if err != nil {
		c.logger.Error("Couldn't create user", zap.ByteString("uri", ctx.Request().RequestURI()), zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "couldn't create user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "user created successfully", "user_id": id})
}
