package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	log        *slog.Logger
	httpServer *http.Server
	port       string
}

func NewServer(log *slog.Logger, port string, handler http.Handler) *Server {
	return &Server{
		log: log,
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
		port: port,
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "server.Run"

	s.log.With(slog.String("op", op), slog.String("port", s.port))

	s.log.Info("Starting http server", slog.String("port", s.port))

	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Stop() error {
	const op = "server.Stop"

	s.log.With(slog.String("op", op))
	s.log.Info("Stopping http server", slog.String("port", s.port))

	return s.httpServer.Shutdown(nil)
}
