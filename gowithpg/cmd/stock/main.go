package main

import (
	"context"
	"gowithpg/config"
	"gowithpg/internal/db/postgres"
	"gowithpg/internal/handler"
	"gowithpg/internal/routes"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// LOAD CONFIG
	cfg := config.MustLoad()

	// LOGGER
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)

	// DATABASE INITIALIZATION
	db, err := postgres.New(cfg)
	if err != nil {

		logger.Error("failed to initialize postgres",
			slog.String("error", err.Error()),
		)

		log.Fatal(err)
	}

	logger.Info("postgres initialized",
		slog.String("env", cfg.Env),
		slog.String("version", "1.0.0"),
	)

	// HANDLER INITIALIZATION
	h := handler.New(db, logger)

	// ROUTES
	router := routes.RegisterRoutes(h)

	// SERVER
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	logger.Info("server startin",
		slog.String("address", cfg.Addr),
	)

	// CHANNEL FOR SHUTDOWN SIGNALS
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// START SERVER
	go func() {

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {

			logger.Error("failed to start serve",
				slog.String("error", err.Error()),
			)

			os.Exit(1)
		}
	}()

	// WAIT FOR SIGNAL
	<-done

	logger.Info("shutting down server")

	// GRACEFUL SHUTDOWN
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		logger.Error("failed to shutdown server",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	logger.Info("server shutdown successfully")
}