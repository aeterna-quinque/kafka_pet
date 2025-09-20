package app

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/controller"
	"kafka-pet/internal/infra/kafka/producer"
	"kafka-pet/internal/infra/logger"
	"kafka-pet/internal/router"
	"kafka-pet/internal/server"
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

	producer, err := producer.NewAsyncProducer(ctx, &config.Kafka)
	if err != nil {
		l.Error("Couldn't create new kafka producer", zap.Error(err))
		return fmt.Errorf("couldn't create new kafka producer: %w", err)
	}
	defer producer.AsyncClose()

	controller := controller.NewController(producer, &config.Kafka)
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
