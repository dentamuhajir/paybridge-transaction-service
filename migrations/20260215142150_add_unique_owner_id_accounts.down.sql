-- +migrate Up

ALTER TABLE accounts
ADD CONSTRAINT accounts_owner_id_unique UNIQUE (owner_id);

-- +migrate Down

ALTER TABLE accounts
DROP CONSTRAINT IF EXISTS accounts_owner_id_unique;
