package producer

import (
	"fmt"
	"kafka-pet/internal/config"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type SyncProducer struct {
	producer sarama.SyncProducer
	logger   *zap.Logger
}

func NewSyncProducer(logger *zap.Logger, cfg *config.Kafka) (*SyncProducer, error) {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Return.Errors = true
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(cfg.Brokers, saramaCfg)
	if err != nil {
		logger.Error("Couldn't create new sync producer", zap.Error(err))
		return nil, fmt.Errorf("couldn't create new sync producer: %w", err)
	}

	return &SyncProducer{
		producer: producer,
		logger:   logger,
	}, nil
}

func (p *SyncProducer) Close() error {
	if err := p.producer.Close(); err != nil {
		p.logger.Error("Couldn't close sync porducer", zap.Error(err))
		return fmt.Errorf("couldn't close sync producer: %w", err)
	}
	return nil
}

func (p *SyncProducer) SendMessage(topic string, key string, message []byte) error {
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(message),
	})
	if err != nil {
		p.logger.Error("Couldn't send message to kafka", zap.String("topic", "users"), zap.Error(err))
		return fmt.Errorf("couldn't send message to kafka: %w", err)
	}

	p.logger.Info(
		"Message has been sent to kafka successfully",
		zap.String("topic", topic),
		zap.String("key", key),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)
	return nil
}
