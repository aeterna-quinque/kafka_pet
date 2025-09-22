package messages

import (
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/domain"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type MessagesConsumerHandler struct {
	messages *domain.Messages
	cfg      *config.MessagesConsumer
	logger   *zap.Logger
}

func NewMessagesConsumerHandler(cfg *config.MessagesConsumer, logger *zap.Logger) *MessagesConsumerHandler {
	return &MessagesConsumerHandler{
		messages: &domain.Messages{
			Messages: make([]string, 0, 10),
		},
		cfg:    cfg,
		logger: logger,
	}
}

func (c *MessagesConsumerHandler) GetTopics() []string {
	return c.cfg.Topics
}

func (c *MessagesConsumerHandler) GetMessages() (*domain.Messages, error) {
	if c.messages == nil {
		return nil, fmt.Errorf("Messages unavailable")
	}

	return c.messages, nil
}

func (c *MessagesConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	c.messages = &domain.Messages{
		Messages: make([]string, 0, 10),
	}
	c.logger.Info("Setup messages in messages consumer")
	return nil
}

func (c *MessagesConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	c.messages = nil
	c.logger.Info("Cleanup messages in messages consumer")
	return nil
}

func (c *MessagesConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		message := fmt.Sprintf("Topic: %s, Key: %s, Value: %s", msg.Topic, msg.Key, msg.Value)
		c.messages.Messages = append(c.messages.Messages, message)
		c.logger.Info("Message consumed", zap.String("msg", message))
		session.MarkMessage(msg, "")
	}
	return nil
}
