package server

import (
	"context"
	"net/http"

	"github.com/NikolayStepanov/PasswordGenerator/internal/config"
)

// Server wraps the HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer creates and returns a new Server instance configured
// with the specified configuration and HTTP handler
func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              cfg.HTTP.Host + ":" + cfg.HTTP.Port,
			Handler:           handler,
			ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
			ReadTimeout:       cfg.HTTP.ReadTimeout,
			WriteTimeout:      cfg.HTTP.WriteTimeout,
			IdleTimeout:       cfg.HTTP.IdleTimeout,
		},
	}
}

// Run starts the HTTP server and listens for incoming requests
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server using the provided context
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
