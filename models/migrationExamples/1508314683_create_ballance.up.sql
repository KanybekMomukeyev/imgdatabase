CREATE TABLE IF NOT EXISTS ballances (
    ballance_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    stock_id BIGINT,
    user_id BIGINT,
    product_id BIGINT,
    customer_id BIGINT,
    supplier_id BIGINT,
    transaction_id BIGINT,
    order_detail_id BIGINT,
    ballance_date BIGINT,
    ballance_type INTEGER,
    value REAL,
    uuid varchar (400),
    comment varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_ballances_idx ON ballances (company_id);
CREATE INDEX IF NOT EXISTS user_id_ballances_idx ON ballances (user_id);
CREATE INDEX IF NOT EXISTS stock_id_ballances_idx ON ballances (stock_id);
--CREATE INDEX IF NOT EXISTS transaction_id_ballances_idx ON ballances (transaction_id);
--CREATE INDEX IF NOT EXISTS order_detail_id_ballances_idx ON ballances (order_detail_id);
