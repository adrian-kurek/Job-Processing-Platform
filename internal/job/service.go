package job

import "log/slog"

type handlerRepository interface{}

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
