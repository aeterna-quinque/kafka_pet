package consumer

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ConsumerGroup struct {
	group  sarama.ConsumerGroup
	logger *zap.Logger
}

func NewConsumerGroup(ctx context.Context, cfg *config.Kafka) (*ConsumerGroup, error) {
	logger := logger.FromContext(ctx)

	saramaCfg := sarama.NewConfig()

	group, err := sarama.NewConsumerGroup(cfg.Brokers, "users1", saramaCfg)
	if err != nil {
		logger.Error("Couldn't create consumer group", zap.Error(err))
		return nil, fmt.Errorf("couldn't create consumer group: %w", err)
	}

	g := &ConsumerGroup{
		group:  group,
		logger: logger,
	}

	return g, nil
}

func (g *ConsumerGroup) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) {
	for {
		select {
		case <-ctx.Done():
			g.logger.Info("Consumer group stopped consuming due to context done")
			return
		default:
			if err := g.group.Consume(ctx, topics, handler); err != nil {
				g.logger.Error("Consumer group couldn't consume message", zap.Error(err))
			}
			g.logger.Info("Consumer group consumed message successfully")
		}
	}
}

func (g *ConsumerGroup) Close() error {
	if err := g.group.Close(); err != nil {
		g.logger.Error("Couldn't close consumer group", zap.Error(err))
		return fmt.Errorf("couldn't close consumer group: %w", err)
	}
	return nil
}
