package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gitgotgon/go-shorter/internal/config"
	apphttp "github.com/gitgotgon/go-shorter/internal/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	router := apphttp.NewRouter()

	server := &http.Server{
		Addr:              ":" + strconv.Itoa(cfg.HTTPPort),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Cancel this context when the process receives SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(
		context.Background(),
		// Ctrl-c, interactive interruption by a user
		os.Interrupt,
		// kill <pid>, tells the process to shut down gracefully, commonly send by OS, service manager, or kill <pid>
		syscall.SIGTERM, 
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("HTTP server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server: %w", err)
		}
		return nil

	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		_ = server.Close()
		return fmt.Errorf("graceful shutdown: %w", err)
	}

	log.Println("HTTP server stopped gracefully")
	return nil
}
