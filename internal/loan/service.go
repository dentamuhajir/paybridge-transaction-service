package loan

import (
	"context"
	"paybridge-transaction-service/internal/loan/entity"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, req LoanAppCreateRequest) (*LoanAppCreateResponse, error)
	Approval(ctx context.Context, req LoanApprovalRequest) (*LoanApprovalResponse, error)
	BatchApproval(ctx context.Context, req []LoanApprovalRequest) (*LoanBatchApprovalResponse, error)
}

type service struct {
	repo Repository
	log  *zap.Logger
}

func NewService(r Repository, log *zap.Logger) Service {
	return &service{repo: r, log: log}
}

func (s *service) BatchApproval(ctx context.Context, req []LoanApprovalRequest) (*LoanBatchApprovalResponse, error) {

	type result struct {
		id     string
		status string
		err    error
	}

	resultsCh := make(chan result, len(req))
	var wg sync.WaitGroup

	for _, r := range req {
		wg.Add(1)

		go func(r LoanApprovalRequest) {
			defer wg.Done()

			resp, err := s.Approval(ctx, r)
			if err != nil {
				resultsCh <- result{
					id:  r.ID.String(),
					err: err,
				}
				return
			}

			resultsCh <- result{
				id:     resp.ID,
				status: resp.Status,
			}
		}(r)
	}

	wg.Wait()
	close(resultsCh)

	response := LoanBatchApprovalResponse{}
	for r := range resultsCh {
		item := LoanBatchApprovalResult{
			ID:     r.id,
			Status: r.status,
		}

		if r.err != nil {
			item.Error = r.err.Error()
			response.FailedCount++
		} else {
			response.SuccessCount++
		}

		response.Results = append(response.Results, item)
	}

	response.TotalLoanUpdate = response.SuccessCount

	return &response, nil
}

func (s *service) Create(ctx context.Context, req LoanAppCreateRequest) (*LoanAppCreateResponse, error) {

	loan := entity.LoanApplication{
		UserID:                  req.UserID,
		ProductID:               req.ProductID,
		Amount:                  req.Amount,
		TenorMonth:              req.TenorMonth,
		InterestType:            req.InterestType,
		AdminFee:                req.AdminFee,
		DisbursementScheduledAt: req.DisbursementScheduledAt,
	}

	result, err := s.repo.Create(ctx, loan)
	if err != nil {
		log.Error(ctx, "error in service", err)
		return nil, err
	}

	return &LoanAppCreateResponse{
		ID:     result.ID.String(),
		Status: result.Status,
	}, nil
}

func (s *service) Approval(ctx context.Context, req LoanApprovalRequest) (*LoanApprovalResponse, error) {
	loan := entity.LoanApplication{
		ID:        req.ID,
		Status:    req.Status,
		UpdatedAt: time.Now(),
	}

	result, err := s.repo.Approval(ctx, loan)

	if err != nil {
		return nil, err
	}

	return &LoanApprovalResponse{
		ID:     result.ID.String(),
		Status: req.Status,
	}, nil

}
