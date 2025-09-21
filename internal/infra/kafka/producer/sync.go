package producer

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

func NewSyncProducer(ctx context.Context, cfg *config.Kafka) (sarama.SyncProducer, error) {
	l := logger.FromContext(ctx)

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Return.Errors = true
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(cfg.Brokers, saramaCfg)
	if err != nil {
		l.Error("Couldn't create new sync producer", zap.Error(err))
		return nil, fmt.Errorf("couldn't create new sync producer: %w", err)
	}

	return producer, nil
}
