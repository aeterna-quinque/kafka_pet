package service

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/infra/logger"
	"strconv"
	"time"

	"go.uber.org/zap"
)

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

	s.asyncProducer.SendMessage(
		s.cfg.Kafka.UsersGetTopic,
		strconv.Itoa(int(id)),
		eventJson,
	)

	return &domain.User{}, nil
}
