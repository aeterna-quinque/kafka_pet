package controller

import (
	"kafka-pet/internal/dto"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (c *Controller) GetMessages(ctx *fiber.Ctx) error {
	messages, err := c.service.GetMessages()
	if err != nil {
		c.logger.Error("Couldn't get messages", zap.ByteString("uri", ctx.Request().RequestURI()))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "Couldn't get messages. Try later"})
	}

	c.logger.Info("Messages", zap.Any("messages", messages))
	resp := dto.MessagesToGetMessagesResponse(messages)
	c.logger.Info("Response", zap.Any("response", resp))

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
