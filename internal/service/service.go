package service

import (
	"kafka-pet/internal/config"

	"github.com/IBM/sarama"
)

type Service struct {
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
	cfg           *config.Config
}

func NewService(syncProducer sarama.SyncProducer, asyncProducer sarama.AsyncProducer, cfg *config.Config) Servicer {
	return &Service{
		syncProducer:  syncProducer,
		asyncProducer: asyncProducer,
		cfg:           cfg,
	}
}
