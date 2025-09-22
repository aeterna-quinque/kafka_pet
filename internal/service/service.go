package service

import (
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/kafka/producer"
	"kafka-pet/internal/messages"
	"kafka-pet/internal/repo/postgres"
)

type Service struct {
	syncProducer     *producer.SyncProducer
	asyncProducer    *producer.AsyncProducer
	cfg              *config.Config
	messagesConsumer *messages.MessagesConsumerHandler
	repo             *postgres.Repository
}

func NewService(
	syncProducer *producer.SyncProducer,
	asyncProducer *producer.AsyncProducer,
	cfg *config.Config,
	statsConsumer *messages.MessagesConsumerHandler,
	repo *postgres.Repository,
) Servicer {
	return &Service{
		syncProducer:     syncProducer,
		asyncProducer:    asyncProducer,
		cfg:              cfg,
		messagesConsumer: statsConsumer,
		repo:             repo,
	}
}
