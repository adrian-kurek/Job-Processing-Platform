package job

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	commonerrors "github.com/adrian-kurek/Job-Processing-Platform/common/errors"
	"github.com/adrian-kurek/Job-Processing-Platform/common/middleware"
	"github.com/adrian-kurek/Job-Processing-Platform/common/request"
)

type handlerService interface {
	Insert(ctx context.Context, job CreateDTO) error
}

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

const CRUDTimeout = 5 * time.Second

func (jh *Handler) handleTimeout(err error, path string) error {
	if errors.Is(err, context.DeadlineExceeded) {
		jh.logger.Info("request timed out", "path", path)
		return commonerrors.RequestTimeout()
	}
	return err
}

func (jh *Handler) Insert(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), CRUDTimeout)
	defer cancel()

	reqData, err := request.ReadBody[CreateDTO](r)
	if err != nil {
		return commonerrors.InvalidJSONFormat()
	}

	err = middleware.ValidateRequestData(reqData)
	if err != nil {
		return err
	}

	err = jh.handlerService.Insert(ctx, *reqData)
	if err != nil {
		return jh.handleTimeout(err, r.URL.Path)
	}
	return nil
}
