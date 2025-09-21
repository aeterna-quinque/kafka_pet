package service

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/dto"
	"kafka-pet/internal/infra/logger"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

const createSubTopic = "create"

func (s *Service) CreateUser(ctx context.Context, req *dto.CreateUserRequest) error {
	l := logger.FromContext(ctx)

	user := req.ToUser()

	event := domain.CreateUserEvent{
		UserId:    user.Id,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: time.Now(),
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		l.Error("Couldn't marshal create user event", zap.Error(err))
		return fmt.Errorf("couldn't marshal create user event: %w", err)
	}

	topic := fmt.Sprintf("%s.%s", s.cfg.Kafka.UsersTopic, createSubTopic)
	key := strconv.Itoa(int(user.Id))

	partition, offset, err := s.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(eventJson),
	})
	if err != nil {
		l.Error("Couldn't send message to kafka", zap.String("topic", "users"), zap.Error(err))
		return fmt.Errorf("couldn't send message to kafka: %w", err)
	}
	l.Info(
		"Message has been sent to kafka successfully",
		zap.String("topic", topic),
		zap.String("key", key),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	l.Info("User created successfully", zap.Uint32("user_id", user.Id))

	return nil
}
