package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ECecillo/GoProm/middleware"
	"github.com/ECecillo/GoProm/routes"
)

type Config struct {
	Host string
	Port string
}

func NewServer(config *Config) http.Handler {
	mux := http.NewServeMux()
	routes.AddRoutes(mux)

	middelwares := middleware.CreateStack(
		middleware.HttpRequestLogger,
		middleware.Prometheus,
	)

	handler := middelwares(mux)

	return handler
}

func Create(ctx context.Context,config *Config) error {
	srv := NewServer(&Config{})
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
