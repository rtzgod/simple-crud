package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/rtzgod/logger"
	"github.com/rtzgod/simple-crud/internal/config"
	"github.com/rtzgod/simple-crud/internal/handler"
	"github.com/rtzgod/simple-crud/internal/repository"
	"github.com/rtzgod/simple-crud/internal/service"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	db, err := repository.NewPostgres(cfg.Postgres.Url)
	if err != nil {
		log.Error("Failed to connect to database", logger.Err(err))
		os.Exit(1)
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handlers := handler.NewHandler(service)

	server := handler.NewServer(log, cfg.HTTP.Port, handlers.InitRoutes())

	go server.MustRun()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sign := <-sigChan

	log.Info("Stopping server", slog.String("signal:", sign.String()))

	if err := server.Stop(); err != nil {
		log.Error("Failed to stop server", logger.Err(err))
	}

	log.Info("Server stopped")
}
