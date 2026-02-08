package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"paybridge-transaction-service/internal/config"
	"paybridge-transaction-service/internal/infra/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, database *pgxpool.Pool, log *logger.Logger) error {
	// Setup router
	e := NewRouter(database, log)

	// Start server in goroutine
	go func() {
		log.Info(context.Background(), "server starting",
			zap.Int("port", cfg.Server.Port),
		)

		if err := e.Start(":" + strconv.Itoa(cfg.Server.Port)); err != nil &&
			err != http.ErrServerClosed {
			log.Error(context.Background(), "server failed", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info(context.Background(), "shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error(context.Background(), "forced shutdown", err)
	} else {
		log.Info(context.Background(), "server shut down gracefully")
	}

	return nil
}
