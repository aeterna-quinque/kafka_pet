package controller

import (
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Greet(ctx *fiber.Ctx) error {
	c.producer.Input() <- &sarama.ProducerMessage{
		Topic: c.kafkaCfg.ServerTopic,
		Key:   sarama.StringEncoder("request"),
		Value: sarama.StringEncoder(ctx.Request().String()),
	}
	defer func() {
		c.producer.Input() <- &sarama.ProducerMessage{
			Topic: c.kafkaCfg.ServerTopic,
			Key:   sarama.StringEncoder("response"),
			Value: sarama.StringEncoder(ctx.Response().String()),
		}
	}()
	return ctx.Status(fiber.StatusOK).SendString("Welcome!")
}
