package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nkchakradhari780/practice9/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// db connection
	// chi router setup
	router := chi.NewRouter()
	// server setup

	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("starting server....")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func(){
		if err := server.ListenAndServe(); err != nil {
			slog.Error("cannot start server", "", err.Error())
		}
	}()

	<- done

	slog.Info("shutting down server.........")

	ctx, stop := context.WithTimeout(context.Background(), 10*time.Second)
	defer stop()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("forcing shutdown server", "error", err)
	}

	slog.Info("server shutdown successfully......")

}