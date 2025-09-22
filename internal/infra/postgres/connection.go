package postgres

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewPostgres(ctx context.Context, cfg *config.Postgres) (*pgxpool.Pool, error) {
	l := logger.FromContext(ctx)

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		l.Error("Couldn't connect to postgres", zap.Error(err))
		return nil, fmt.Errorf("couldn't connect to postgres: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		l.Error("Couldn't ping postgres", zap.Error(err))
		return nil, fmt.Errorf("couldn't ping postgres: %w", err)
	}

	return db, nil
}
