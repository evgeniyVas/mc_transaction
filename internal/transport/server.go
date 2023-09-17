package transport

import (
	"context"
	"github.com/mc_transaction/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.HttpServer.Port,
			Handler:      handler,
			ReadTimeout:  cfg.HttpServer.ReadTimeout,
			WriteTimeout: cfg.HttpServer.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
