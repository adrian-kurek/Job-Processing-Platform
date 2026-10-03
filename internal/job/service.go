package job

import (
	"context"
	"log/slog"
)

type handlerRepository interface {
	Insert(ctx context.Context, job CreateDTO) error
}

type Service struct {
	handlerRepository handlerRepository
	logger            *slog.Logger
}

func NewService(handlerRepository handlerRepository, logger *slog.Logger) *Service {
	return &Service{
		handlerRepository: handlerRepository,
		logger:            logger,
	}
}

func (js *Service) Insert(ctx context.Context, job CreateDTO) error {
	return js.handlerRepository.Insert(ctx, job)
}
