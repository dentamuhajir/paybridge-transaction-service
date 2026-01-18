package account

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateAccountWithInitialBalances(ctx context.Context, userID uuid.UUID) error
}

type service struct {
	repo repository
	log  *zap.Logger
}

func NewService(repo repository, log *zap.Logger) Service {
	return &service{repo, log}
}

func (s *service) CreateAccountWithInitialBalances(ctx context.Context, userID uuid.UUID) error {
	return s.repo.CreateAccountWithBalance(ctx, userID)
}
