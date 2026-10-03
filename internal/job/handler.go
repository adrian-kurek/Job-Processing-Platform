package job

import "log/slog"

type handlerService interface{}

type Handler struct {
	handlerService handlerService
	logger         *slog.Logger
}

func NewHandler(handlerService handlerService, logger *slog.Logger) *Handler {
	return &Handler{
		handlerService: handlerService,
		logger:         logger,
	}
}


