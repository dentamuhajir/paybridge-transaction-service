package loan

import (
	"context"
	"errors"
	"fmt"
	"paybridge-transaction-service/internal/loan/entity"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository interface {
	Create(ctx context.Context, loan entity.LoanApplication) (entity.LoanApplication, error)
	Approval(ctx context.Context, loan entity.LoanApplication) (entity.LoanApplication, error)
}

var ErrLoanNotPendingOrNotFound = errors.New("loan not pending or not found")

type repository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewRepository(db *pgxpool.Pool, log *zap.Logger) Repository {
	return &repository{db: db, log: log}
}

func (r *repository) Create(ctx context.Context, loan entity.LoanApplication) (entity.LoanApplication, error) {
	now := time.Now()
	loan.CreatedAt = now
	loan.UpdatedAt = now
	loan.Status = "PENDING"

	query := `
		INSERT INTO loan_applications 
			(user_id, product_id, amount, tenor_month, interest_type, admin_fee, status, disbursement_scheduled_at, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id, user_id, product_id, amount, tenor_month, interest_type, admin_fee, status, disbursement_scheduled_at, disbursed_at, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		loan.UserID,
		loan.ProductID,
		loan.Amount,
		loan.TenorMonth,
		loan.InterestType,
		loan.AdminFee,
		loan.Status,
		loan.DisbursementScheduledAt,
		loan.CreatedAt,
		loan.UpdatedAt,
	).Scan(
		&loan.ID,
		&loan.UserID,
		&loan.ProductID,
		&loan.Amount,
		&loan.TenorMonth,
		&loan.InterestType,
		&loan.AdminFee,
		&loan.Status,
		&loan.DisbursementScheduledAt,
		&loan.DisbursedAt,
		&loan.CreatedAt,
		&loan.UpdatedAt,
	)

	if err != nil {
		r.log.Error("failed to create loan application", zap.Error(err))
		return entity.LoanApplication{}, err
	}

	return loan, nil
}

func (r *repository) Approval(ctx context.Context, loan entity.LoanApplication) (entity.LoanApplication, error) {
	now := time.Now()
	loan.CreatedAt = now

	fmt.Printf("%+v\n", loan)

	query := `
		UPDATE loan_applications SET status = $2, updated_at = $3 WHERE id = $1 AND status = 'PENDING'
		RETURNING id, status, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		loan.ID,
		loan.Status,
		loan.UpdatedAt,
	).Scan(
		&loan.ID,
		&loan.Status,
		&loan.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			r.log.Error("Err Loan Not Pending Or Not Found ", zap.Error(err))
			return entity.LoanApplication{}, ErrLoanNotPendingOrNotFound
		}
		r.log.Error("failed to approve loan application ", zap.Error(err))
		return entity.LoanApplication{}, err
	}

	return loan, nil
}
