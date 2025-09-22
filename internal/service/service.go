package service

import (
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/kafka/producer"
	"kafka-pet/internal/messages"
)

type Service struct {
	syncProducer     *producer.SyncProducer
	asyncProducer    *producer.AsyncProducer
	cfg              *config.Config
	messagesConsumer *messages.MessagesConsumerHandler
}

func NewService(syncProducer *producer.SyncProducer, asyncProducer *producer.AsyncProducer, cfg *config.Config, statsConsumer *messages.MessagesConsumerHandler) Servicer {
	return &Service{
		syncProducer:     syncProducer,
		asyncProducer:    asyncProducer,
		cfg:              cfg,
		messagesConsumer: statsConsumer,
	}
}
