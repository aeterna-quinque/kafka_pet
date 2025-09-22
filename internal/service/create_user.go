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

func (s *Service) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (uint32, error) {
	l := logger.FromContext(ctx)

	user := req.ToUser()

	id, err := s.repo.AddUser(ctx, user)
	if err != nil {
		l.Error("Couldn't add user to db", zap.Error(err))
		return 0, fmt.Errorf("couldn't add user to db: %w", err)
	}

	event := domain.CreateUserEvent{
		UserId:    id,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: time.Now(),
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		l.Error("Couldn't marshal create user event", zap.Error(err))
		return 0, fmt.Errorf("couldn't marshal create user event: %w", err)
	}

	if err = s.syncProducer.SendMessage(
		s.cfg.Kafka.UsersCreateTopic,
		strconv.Itoa(int(id)),
		eventJson,
	); err != nil {
		l.Error("Couldn't send message to kafka", zap.Error(err))
		return 0, fmt.Errorf("couldn't send message to kafka: %w", err)
	}

	l.Info("User created successfully", zap.Uint32("id", id))
	return id, nil
}
