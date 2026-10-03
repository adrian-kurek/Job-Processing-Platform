package server

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultReadTimeout  = 50 * time.Second
	defaultWriteTimeout = 50 * time.Second
	defaultIdleTimeout  = 30 * time.Second
)

type DependencyConfig struct {
	port string
}

func NewDependencyConfig(port string) *DependencyConfig {
	return &DependencyConfig{
		port: port,
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
