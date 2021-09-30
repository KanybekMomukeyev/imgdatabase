CREATE TABLE IF NOT EXISTS stocks (
    stock_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    stock_uuid VARCHAR (300),
    stock_name VARCHAR (300),
    stock_currency VARCHAR (300),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_stocks_idx ON stocks (company_id);