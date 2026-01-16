package entity

import (
	"time"

	"github.com/google/uuid"
)

type LoanApplication struct {
	ID                      uuid.UUID  `db:"id"`
	UserID                  uuid.UUID  `db:"user_id"`
	ProductID               uuid.UUID  `db:"product_id"`
	Amount                  int64      `db:"amount"`
	TenorMonth              int        `db:"tenor_month"`
	InterestType            string     `db:"interest_type"` // FLAT / ANNUITY
	AdminFee                int64      `db:"admin_fee"`
	Status                  string     `db:"status"`       // PENDING, APPROVED, REJECTED, DISBURSED, COMPLETED
	DisbursedAt             *time.Time `db:"disbursed_at"` // Nullable, filled when disbursed so we add pointer
	DisbursementScheduledAt time.Time  `db:"disbursement_scheduled_at"`
	CreatedAt               time.Time  `db:"created_at"`
	UpdatedAt               time.Time  `db:"updated_at"`
}
