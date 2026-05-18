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
	// "gowithpg/internal/db"
)
func main(){
	cfg:=config.MustLoad()

	//database initialisation 
	db,err:=postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Postgres initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	
	//initialising handler 
	h:=&handler.Handler{
		DB: db,
	}
	router:=routes.RegisterRoutes(h)
	server:=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.Addr))
	slog.Info("server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}

