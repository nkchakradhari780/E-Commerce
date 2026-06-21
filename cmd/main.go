package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nkchakradhari780/practice9/internal/config"
	"github.com/nkchakradhari780/practice9/internal/storage/postgres"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// db connection
	pg, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Db.Close()
	// chi router setup
	router := chi.NewRouter()
	// server setup

	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("starting server....")

	go func(){
		if err := server.ListenAndServe(); err != nil  && err != http.ErrServerClosed{
			slog.Error("cannot start server", "error", err)
		}
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	<- ctx.Done()

	slog.Info("shutting down server.........")

	shutdownContext, cancle := context.WithTimeout(
		context.Background(), 
		10*time.Second,
	)
	defer cancle()

	if err := server.Shutdown(shutdownContext); err != nil {
		slog.Error("forcing shutdown server", "error", err)
	}

	slog.Info("server shutdown successfully......")

}