package config

import (
	"context"
	"fmt"
	"kafka-pet/internal/infra/logger"

	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
)

type Config struct {
	Server Server `envPrefix:"SERVER_"`
	Kafka  Kafka  `envPrefix:"KAFKA_"`
}

func Load(ctx context.Context) (*Config, error) {
	l := logger.FromContext(ctx)

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		l.Error("Couldn't parse env", zap.Error(err))
		return nil, fmt.Errorf("couldn't parse env: %w", err)
	}

	l.Info("Config loaded successfully")
	return &cfg, nil
}
