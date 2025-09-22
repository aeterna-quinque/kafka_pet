package service

import (
	"fmt"
	"kafka-pet/internal/domain"
)

func (s *Service) GetMessages() (*domain.Messages, error) {
	messages, err := s.messagesConsumer.GetMessages()
	if err != nil {
		return nil, fmt.Errorf("Couldn't retrieve messages: %w", err)
	}

	return messages, nil
}
