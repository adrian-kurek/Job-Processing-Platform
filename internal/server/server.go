package server

import (
	"context"
	"net/http"
	"time"

	"github.com/adrian-kurek/Job-Processing-Platform/internal/job"
)

const (
	defaultReadTimeout  = 50 * time.Second
	defaultWriteTimeout = 50 * time.Second
	defaultIdleTimeout  = 30 * time.Second
)

type DependencyConfig struct {
	port       string
	jobHandler job.Handler
}

func NewDependencyConfig(port string, jobHandler job.Handler) *DependencyConfig {
	return &DependencyConfig{
		port:       port,
		jobHandler: jobHandler,
	}
}

type HTTP struct {
	config *DependencyConfig
	server *http.Server
	router *http.ServeMux
}

func NewHTTP(config *DependencyConfig) *HTTP {
	return &HTTP{
		config: config,
		router: http.NewServeMux(),
	}
}

func (h *HTTP) setupRoutes() {
	jobRoute := job.NewRoute(&h.config.jobHandler)
	jobRoute.SetupRoutes(h.router)
}

func (h *HTTP) Start() error {
	h.setupRoutes()
	h.server = &http.Server{
		Addr:         ":" + h.config.port,
		Handler:      h.router,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}
	return h.server.ListenAndServe()
}

func (h *HTTP) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
