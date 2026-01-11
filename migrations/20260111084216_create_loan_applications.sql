-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS loan_applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    product_id UUID NOT NULL UNIQUE,
    amount BIGINT NOT NULL DEFAULT 0,
    tenor_month INT NOT NULL,
    interest_type VARCHAR(20) NOT NULL, -- FLAT / ANNUTY
    admin_fee BIGINT NOT NULL DEFAULT 0,
    status VARCHAR(20) NOT NULL, -- PENDING, APPROVED, REJECTED, DISBURSED, COMPLETED
    disbursed BOOLEAN DEFAULT false,
    disbursed_at TIMESTAMPTZ NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- +migrate Down
DROP TABLE IF EXISTS loan_applications;
