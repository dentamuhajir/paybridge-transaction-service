-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS account_balances (
    account_id UUID NOT NULL,
    balance_type_id INT NOT NULL,
    amount BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (account_id, balance_type_id),
    CONSTRAINT fk_account_balance_account
        FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_account_balance_type
        FOREIGN KEY (balance_type_id) REFERENCES balance_types(id)
);

-- +migrate Down
DROP TABLE IF EXISTS account_balances;