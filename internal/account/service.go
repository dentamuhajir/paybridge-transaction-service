package account

import (
	"context"
	"paybridge-transaction-service/internal/infra/logger"

	"github.com/google/uuid"
)

type Service interface {
	CreateAccountWithInitialBalances(ctx context.Context, userID uuid.UUID) error
	GetAccount(ctx context.Context, userID uuid.UUID) (Account, error)
}

type service struct {
	repo Repository
	log  *logger.Logger
}

func NewService(repo Repository, log *logger.Logger) Service {
	return &service{repo, log}
}

func (s *service) CreateAccountWithInitialBalances(ctx context.Context, userID uuid.UUID) error {
	return s.repo.CreateAccountWithBalance(ctx, userID)
}

func (s *service) GetAccount(ctx context.Context, userID uuid.UUID) (Account, error) {
	if userID == uuid.Nil {
		return Account{}, ErrInvalidUserID
	}

	account, err := s.repo.GetAccount(ctx, userID)

	if err != nil {
		return Account{}, err
	}

	if account.Status != StatusActive {
		return Account{}, ErrAccountInactive
	}

	return account, err
}
