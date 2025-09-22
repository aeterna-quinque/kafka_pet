package controller

import (
	"errors"
	"kafka-pet/internal/dto"
	"kafka-pet/internal/service"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (c *Controller) GetUser(ctx *fiber.Ctx) error {
	idStr := ctx.Query("id")
	if idStr == "" {
		c.logger.Error("No ID provided", zap.ByteString("uri", ctx.Request().RequestURI()))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.ErrBadRequest, "message": "No ID provided"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.logger.Error("Couldn't parse ID", zap.ByteString("uri", ctx.Request().RequestURI()), zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.ErrBadRequest, "message": "Couldn't parse ID"})
	}

	if id <= 0 || id > math.MaxUint32 {
		c.logger.Error("ID must be in range [1;4294967295]", zap.ByteString("uri", ctx.Request().RequestURI()))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.ErrBadRequest, "message": "ID must be in range [1;4294967295]"})
	}

	user, err := c.service.GetUser(c.ctx, uint32(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.logger.Info("User not found", zap.ByteString("uri", ctx.Request().RequestURI()), zap.Int("id", id), zap.Error(err))
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "User not found"})
		}
		c.logger.Error("Couldn't get user", zap.ByteString("uri", ctx.Request().RequestURI()), zap.Int("id", id), zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.ErrInternalServerError, "message": "Couldn't get user"})
	}

	resp := dto.GetUserResponseFromUser(user)

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
