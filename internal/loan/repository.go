package loan

import (
	"context"
	"paybridge-transaction-service/internal/domain/loan/entity"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository interface {
	Create(ctx context.Context, loan entity.LoanApplication) (entity.LoanApplication, error)
}

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
	loan.Disbursed = false

	query := `
		INSERT INTO loan_applications 
			(user_id, product_id, amount, tenor_month, interest_type, admin_fee, status, disbursed, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id, user_id, product_id, amount, tenor_month, interest_type, admin_fee, status, disbursed, disbursed_at, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		loan.UserID,
		loan.ProductID,
		loan.Amount,
		loan.TenorMonth,
		loan.InterestType,
		loan.AdminFee,
		loan.Status,
		loan.Disbursed,
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
		&loan.Disbursed,
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
