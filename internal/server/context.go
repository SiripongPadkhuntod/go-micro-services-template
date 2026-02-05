package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	servicesconstant "go-micro-services-template/internal/constant"
)

func RunServerWithContext(ctx context.Context, h http.Handler) error {
	// cfg := property.Get()
	port := "8080"
	// if cfg.Server.Port != "" {
	// 	port = cfg.Server.Port
	// }

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h,
	}

	errChan := make(chan error, 1)

	// Run server in background
	go func() {
		log.Printf("Start server on port %s at %s\n", port, time.Now().Local().Format(servicesconstant.TIME_LAYOUT_DATETIME_FORMAT))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Wait for signal or error
	select {
	case err := <-errChan:
		log.Printf("Error from server: %v", err)
		return err
	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), servicesconstant.DefaultShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)
		return err
	}

	log.Println("Server stopped cleanly")
	return nil
}
