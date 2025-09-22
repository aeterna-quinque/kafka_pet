package service

import (
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/kafka/producer"
)

type Service struct {
	syncProducer  *producer.SyncProducer
	asyncProducer *producer.AsyncProducer
	cfg           *config.Config
}

func NewService(syncProducer *producer.SyncProducer, asyncProducer *producer.AsyncProducer, cfg *config.Config) Servicer {
	return &Service{
		syncProducer:  syncProducer,
		asyncProducer: asyncProducer,
		cfg:           cfg,
	}
}
