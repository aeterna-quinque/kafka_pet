package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/infra/logger"
	"kafka-pet/internal/repo/postgres"
	"strconv"
	"time"

	"go.uber.org/zap"
)

var ErrUserNotFound = errors.New("User not found")

func (s *Service) GetUser(ctx context.Context, id uint32) (*domain.User, error) {
	l := logger.FromContext(ctx)

	userModel, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		l.Error("Couldn't get user from db", zap.Error(err))
		return nil, fmt.Errorf("couldn't get user from db: %w", err)
	}

	if userModel == nil {
		l.Info("User not found", zap.Uint32("id", id))
		return nil, nil
	}

	user := domain.UserEntityFromUserModel(userModel)

	event := domain.GetUserEvent{
		UserId:      user.Id,
		RequestedAt: time.Now(),
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		l.Error("Couldn't marshal get user event", zap.Error(err))
		return nil, fmt.Errorf("couldn't marshal get user event: %w", err)
	}

	s.asyncProducer.SendMessage(
		s.cfg.Kafka.UsersGetTopic,
		strconv.Itoa(int(user.Id)),
		eventJson,
	)

	return user, nil
}
