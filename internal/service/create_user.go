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

	if err = s.syncProducer.SendMessage(
		fmt.Sprintf("%s.%s", s.cfg.Kafka.UsersTopic, createSubTopic),
		strconv.Itoa(int(user.Id)),
		eventJson,
	); err != nil {
		l.Error("Couldn't send message to kafka", zap.Error(err))
		return fmt.Errorf("couldn't send message to kafka: %w", err)
	}

	l.Info("User created successfully", zap.Uint32("user_id", user.Id))
	return nil
}
