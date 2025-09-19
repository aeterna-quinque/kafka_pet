package server

import (
	"context"
	"fmt"
	"kafka-pet/internal/config"
	"kafka-pet/internal/infra/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	server *fiber.App
	cfg    *config.Server
}

func NewServer(cfg *config.Server, router *fiber.App) *Server {
	s := &Server{
		server: fiber.New(),
		cfg:    cfg,
	}

	s.server.Mount("/", router)

	return s
}

func (s *Server) Run(ctx context.Context) error {
	l := logger.FromContext(ctx)

	l.Info("Server is listening...")
	if err := s.server.Listen(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)); err != nil {
		l.Error("Server couldn't start listening", zap.Error(err))
		return fmt.Errorf("server couldn't start listening: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	l := logger.FromContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, s.cfg.ShutdownTimeout)
	defer cancel()

	if err := s.server.ShutdownWithContext(ctx); err != nil {
		l.Error("Server couldn't shutdown gracefully", zap.Error(err))
		return fmt.Errorf("server couldn't shutdown gracefully: %w", err)
	}

	return nil
}
