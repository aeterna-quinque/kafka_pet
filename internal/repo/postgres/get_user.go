package postgres

import (
	"context"
	"errors"
	"fmt"
	"kafka-pet/internal/infra/logger"
	"kafka-pet/internal/models"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

const getUser = `SELECT id, name, age FROM users WHERE id=$1;`

var ErrUserNotFound = errors.New("User not found")

func (r *Repository) GetUser(ctx context.Context, id uint32) (*models.User, error) {
	l := logger.FromContext(ctx)

	row := r.db.QueryRow(ctx, getUser, id)
	var user models.User
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			l.Info("No rows found for requested id", zap.Uint32("id", id))
			return nil, ErrUserNotFound
		}
		l.Error("Couldn't execute insert query", zap.Error(err))
		return nil, fmt.Errorf("couldn't execute insert query: %w", err)
	}

	l.Info("User acquired successfully", zap.Uint32("id", id))
	return &user, nil
}
