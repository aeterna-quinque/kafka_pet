package controller

import (
	"kafka-pet/internal/config"

	"github.com/IBM/sarama"
)

type Controller struct {
	producer sarama.AsyncProducer
	kafkaCfg *config.Kafka
}

func NewController(producer sarama.AsyncProducer, cfg *config.Kafka) *Controller {
	return &Controller{
		producer: producer,
		kafkaCfg: cfg,
	}
}
