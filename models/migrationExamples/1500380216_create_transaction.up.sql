CREATE TABLE IF NOT EXISTS transactions (
    transaction_id BIGSERIAL PRIMARY KEY NOT NULL,
    transaction_date BIGINT,
    company_id BIGINT,
    stock_id BIGINT,
    is_last_transaction INTEGER,
    transaction_type INTEGER,
    money_amount REAL,
    order_id BIGINT,
    customer_id BIGINT,
    supplier_id BIGINT,
  	user_id BIGINT,
  	comment varchar (500),
  	uuid varchar (400),
  	ballance_amount REAL,
  	updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_transactions_idx ON transactions (company_id);
--CREATE INDEX IF NOT EXISTS stock_id_transactions_idx ON transactions (stock_id);
CREATE INDEX IF NOT EXISTS order_id_transactions_idx ON transactions (order_id);
--CREATE INDEX IF NOT EXISTS customer_id_transactions_idx ON transactions (customer_id);
--CREATE INDEX IF NOT EXISTS supplier_id_transactions_idx ON transactions (supplier_id);
--CREATE INDEX IF NOT EXISTS user_id_transactions_idx ON transactions (user_id);