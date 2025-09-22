package controller

import (
	"context"
	"kafka-pet/internal/infra/logger"
	"kafka-pet/internal/service"

	"go.uber.org/zap"
)

type Controller struct {
	service service.Servicer
	ctx     context.Context
	logger  *zap.Logger
}

func NewController(ctx context.Context, service service.Servicer) *Controller {
	return &Controller{
		service: service,
		ctx:     ctx,
		logger:  logger.FromContext(ctx),
	}
}
