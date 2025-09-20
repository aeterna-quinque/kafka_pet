package producer

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

func NewAsyncProducer(ctx context.Context, cfg *config.Kafka) (sarama.AsyncProducer, error) {
	l := logger.FromContext(ctx)

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, sarama.NewConfig())
	if err != nil {
		l.Error("Couldn't create new async producer", zap.Error(err))
		return nil, fmt.Errorf("couldn't create new async producer: %w", err)
	}

	return producer, nil
}
