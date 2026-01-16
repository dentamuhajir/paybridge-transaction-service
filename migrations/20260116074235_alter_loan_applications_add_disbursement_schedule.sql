-- +migrate Up

ALTER TABLE loan_applications
ADD COLUMN disbursement_scheduled_at TIMESTAMPTZ NULL;

-- remove disbursed
ALTER TABLE loan_applications
DROP COLUMN IF EXISTS disbursed;


-- +migrate Down
ALTER TABLE loan_applications
DROP COLUMN IF EXISTS disbursement_scheduled_at;


-- later add index for better performance cron
-- -- +migrate Up
-- CREATE INDEX CONCURRENTLY idx_loan_disbursement_schedule
-- ON loan_applications (status, disbursement_scheduled_at);

-- -- +migrate Down
-- DROP INDEX CONCURRENTLY IF EXISTS idx_loan_disbursement_schedule;
