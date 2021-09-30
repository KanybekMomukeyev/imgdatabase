CREATE TABLE IF NOT EXISTS accounts (
    account_id BIGSERIAL PRIMARY KEY NOT NULL,
    customer_id BIGINT,
    supplier_id BIGINT,
    balance REAL
);