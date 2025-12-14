package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"paybridge-transaction-service/internal/config"

	"strconv"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, database *pgxpool.Pool, log *zap.Logger) error {
	// Setup router
	e := NewRouter(database, log)

	// Start server in goroutine
	go func() {
		log.Info("server starting",
			zap.Int("port", cfg.Server.Port),
		)
		if err := e.Start(":" + strconv.Itoa(cfg.Server.Port)); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal("server failed", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Info("Forced shutdown: " + err.Error())
	} else {
		log.Info("Server shut down gracefully")
	}

	return nil
}
