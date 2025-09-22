package producer

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type AsyncProducer struct {
	producer sarama.AsyncProducer
	logger   *zap.Logger
}

func NewAsyncProducer(ctx context.Context, cfg *config.Kafka) (*AsyncProducer, error) {
	l := logger.FromContext(ctx)

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Return.Errors = true
	saramaCfg.Producer.RequiredAcks = sarama.WaitForLocal

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, saramaCfg)
	if err != nil {
		l.Error("Couldn't create new async producer", zap.Error(err))
		return nil, fmt.Errorf("couldn't create new async producer: %w", err)
	}

	p := &AsyncProducer{
		producer: producer,
		logger:   l,
	}

	go p.parseChannels(ctx)

	return p, nil
}

func (p *AsyncProducer) SendMessage(topic string, key string, message []byte) {
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	p.logger.Info(
		"Message has been put in queue for sending to kafka",
		zap.String("topic", topic),
		zap.String("key", key),
	)
}

func (p *AsyncProducer) Close() {
	p.producer.AsyncClose()
}

func (p *AsyncProducer) parseChannels(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			p.logger.Info("Async producer stopped parsing channels due to context done")
			return
		case err := <-p.producer.Errors():
			p.logger.Error("Async producer couldn't send message", zap.Error(err))
		case msg := <-p.producer.Successes():
			key, err := msg.Key.Encode()
			if err != nil {
				p.logger.Error("Async producer couldn't decode key from successes chan message", zap.Error(err))
				continue
			}
			p.logger.Info("Async producer sent message successfully", zap.String("topic", msg.Topic), zap.ByteString("key", key))
		}
	}
}
