package postgres

import (
	"context"
	"fmt"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/infra/logger"

	"go.uber.org/zap"
)

const addUser = `INSERT INTO users(name, age) VALUES ($1, $2) RETURNING (id);`

func (r *Repository) AddUser(ctx context.Context, user *domain.User) (uint32, error) {
	l := logger.FromContext(ctx)

	model := user.ToUserModel()

	row := r.db.QueryRow(ctx, addUser, model.Name, model.Age)
	var id uint32
	if err := row.Scan(&id); err != nil {
		l.Error("Couldn't execute insert query", zap.Error(err))
		return 0, fmt.Errorf("couldn't execute insert query: %w", err)
	}

	l.Info("User added successfully", zap.Uint32("id", id))
	return id, nil
}
