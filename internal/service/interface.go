package service

import (
	"context"
	"kafka-pet/internal/domain"
	"kafka-pet/internal/dto"
)

type Servicer interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) error
	GetUser(ctx context.Context, id uint32) (*domain.User, error)
}
