package app

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/controller"
	"kafka-pet/internal/infra/kafka/consumer"
	"kafka-pet/internal/infra/kafka/producer"
	"kafka-pet/internal/infra/logger"
	"kafka-pet/internal/messages"
	"kafka-pet/internal/router"
	"kafka-pet/internal/server"
	"kafka-pet/internal/service"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	l := logger.NewLogger(logLevel, os.Stdout)
	ctx = logger.WithLogger(ctx, l)

	config, err := config.Load(ctx)
	if err != nil {
		l.Error("Couldn't load config", zap.Error(err))
		return fmt.Errorf("couldn't load config: %w", err)
	}

	asyncProducer, err := producer.NewAsyncProducer(ctx, &config.Kafka)
	if err != nil {
		l.Error("Couldn't create new kafka async producer", zap.Error(err))
		return fmt.Errorf("couldn't create new kafka async producer: %w", err)
	}
	defer asyncProducer.Close()

	syncProducer, err := producer.NewSyncProducer(l, &config.Kafka)
	if err != nil {
		l.Error("Couldn't create new kafka sync producer", zap.Error(err))
		return fmt.Errorf("couldn't create new kafka sync producer: %w", err)
	}
	defer func() {
		if err = syncProducer.Close(); err != nil {
			l.Error("Couldn't close kafka sync producer", zap.Error(err))
		}
	}()

	consumerGroup, err := consumer.NewConsumerGroup(ctx, &config.Kafka)
	if err != nil {
		l.Error("Couldn't create new kafka consumer group", zap.Error(err))
		return fmt.Errorf("couldn't create new kafka consumer group: %w", err)
	}
	defer func() {
		if err = consumerGroup.Close(); err != nil {
			l.Error("Couldn't close kafka consumer group", zap.Error(err))
		}
	}()

	statsConsumer := messages.NewMessagesConsumerHandler(&config.MessagesConsumer, l)
	go consumerGroup.Consume(ctx, statsConsumer.GetTopics(), statsConsumer)

	service := service.NewService(syncProducer, asyncProducer, config, statsConsumer)
	controller := controller.NewController(ctx, service)
	router := router.NewRouter(controller)
	server := server.NewServer(&config.Server, router)

	serverErr := make(chan error, 1)
	go func() {
		if err := server.Run(ctx); err != nil {
			l.Error("Server error occured", zap.Error(err))
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		l.Info("Stopping application due to a server error...", zap.Error(err))
		return fmt.Errorf("server error occured: %w", err)

	case <-ctx.Done():
		l.Info("Stopping application by context...")

		ctx := logger.WithLogger(context.Background(), l)
		if err = server.Stop(ctx); err != nil {
			l.Error("Couldn't stop server", zap.Error(err))
			return fmt.Errorf("couldn't stop server: %w", err)
		}

		l.Info("Application stopped successfully")
		return nil
	}
}
