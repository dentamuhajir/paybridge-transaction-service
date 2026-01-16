-- +migrate Up

-- 1. Drop UNIQUE constraints
ALTER TABLE loan_applications
DROP CONSTRAINT IF EXISTS loan_applications_user_id_key;

ALTER TABLE loan_applications
DROP CONSTRAINT IF EXISTS loan_applications_product_id_key;

-- 2. Add normal indexes instead
CREATE INDEX IF NOT EXISTS idx_loan_applications_user_id
ON loan_applications (user_id);

CREATE INDEX IF NOT EXISTS idx_loan_applications_product_id
ON loan_applications (product_id);

-- +migrate Down

-- Remove indexes
DROP INDEX IF EXISTS idx_loan_applications_user_id;
DROP INDEX IF EXISTS idx_loan_applications_product_id;

-- Restore UNIQUE constraints (rollback)
ALTER TABLE loan_applications
ADD CONSTRAINT loan_applications_user_id_key UNIQUE (user_id);

ALTER TABLE loan_applications
ADD CONSTRAINT loan_applications_product_id_key UNIQUE (product_id);
