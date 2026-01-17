-- +migrate Up

INSERT INTO balance_types (id, code, description)
VALUES
  (1, 'CASH', 'User available balance'),
  (2, 'LOAN_PRINCIPAL', 'Outstanding loan principal'),
  (3, 'LOAN_INTEREST', 'Outstanding loan interest'),
  (4, 'FEE', 'Accrued or collected fees'),
  (5, 'ESCROW', 'Funds held temporarily and not yet settled'),
  (6, 'RESERVE', 'Reserved funds not available for spending')
ON CONFLICT (code) DO NOTHING;
