package app

import (
	"paybridge-transaction-service/internal/config"
	"paybridge-transaction-service/internal/infra/logger"
	"paybridge-transaction-service/internal/infra/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Container struct {
	Cfg    *config.Config
	DB     *pgxpool.Pool
	Logger *zap.Logger
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := postgres.NewPostgres(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	log, err := logger.New()
	if err != nil {
		return nil, err
	}

	return &Container{
		Cfg:    cfg,
		DB:     db,
		Logger: log,
	}, nil
}
