package wallet

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	CreateWallet(ctx context.Context, req CreateWalletReq) (*CreateWalletResp, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Create Wallet
// @Summary Create wallet
// @Description Create a new wallet for a user
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body CreateWalletReq true "Wallet creation payload"
// @Success 200 {object} response.SwaggerSuccessResponse
// @Failure 400 {object} response.SwaggerErrorResponse
// @Failure 500 {object} response.SwaggerErrorResponse
// @Security InternalTokenAuth
// @Router /wallet [post]
func (s *service) CreateWallet(ctx context.Context, req CreateWalletReq) (*CreateWalletResp, error) {

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	wallet := &Wallet{
		UserID:   userUUID,
		Balance:  0,
		Currency: req.Currency,
		Status:   "ACTIVE",
	}

	result, err := s.repo.Create(ctx, wallet)
	if err != nil {
		return nil, err
	}

	return &CreateWalletResp{
		ID:      result.ID.String(),
		UserID:  result.UserID.String(),
		Balance: result.Balance,
		Status:  result.Status,
	}, nil

}
