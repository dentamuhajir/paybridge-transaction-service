package app

import (
	"paybridge-transaction-service/internal/account"
	"paybridge-transaction-service/internal/config"
	"paybridge-transaction-service/internal/infra/logger"
	"paybridge-transaction-service/internal/infra/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	AccountService account.Service
}

type Container struct {
	Cfg     *config.Config
	DB      *pgxpool.Pool
	Logger  *logger.Logger
	Service *Service
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

	// walletRepo := wallet.NewRepository(db, log)
	// walletSvc := wallet.NewService(walletRepo, log)
	accountRepo := account.NewRepository(db, log)
	accountSvc := account.NewService(accountRepo, log)

	return &Container{
		Cfg:    cfg,
		DB:     db,
		Logger: log,
		Service: &Service{
			AccountService: accountSvc,
		},
	}, nil
}
