package service

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/infra/logger"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

const getSubTopic = "get"

func (s *Service) GetUser(ctx context.Context, id uint32) (*domain.User, error) {
	l := logger.FromContext(ctx)

	event := domain.GetUserEvent{
		UserId:      id,
		RequestedAt: time.Now(),
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		l.Error("Couldn't marshal get user event", zap.Error(err))
		return nil, fmt.Errorf("couldn't marshal get user event: %w", err)
	}

	topic := fmt.Sprintf("%s.%s", s.cfg.Kafka.UsersTopic, getSubTopic)
	key := strconv.Itoa(int(id))

	s.asyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(eventJson),
	}
	l.Info(
		"Message has been put in queue for sending to kafka",
		zap.String("topic", topic),
		zap.String("key", key),
	)

	return &domain.User{}, nil
}
