CREATE TABLE IF NOT EXISTS orders (
    order_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    order_document INTEGER,
    money_movement INTEGER,
    billing_no varchar (400),

    user_id BIGINT,
    from_stock_id BIGINT,
    to_stock_id BIGINT,

    customer_id BIGINT,
    supplier_id BIGINT,
    order_date BIGINT,
    payment_id BIGINT,

    error_msg varchar (400),
    uuid varchar (400),
    comment varchar (400),
    is_deleted INTEGER,
    is_money_for_debt INTEGER,
    is_editted INTEGER,
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_orders_idx ON orders (company_id);
--CREATE INDEX IF NOT EXISTS customer_id_orders_idx ON orders (customer_id);
CREATE INDEX IF NOT EXISTS user_id_orders_idx ON orders (user_id);
--CREATE INDEX IF NOT EXISTS supplier_id_orders_idx ON orders (supplier_id);
--CREATE INDEX IF NOT EXISTS payment_id_orders_idx ON orders (payment_id);
CREATE INDEX IF NOT EXISTS from_stock_id_orders_idx ON orders (from_stock_id);
CREATE INDEX IF NOT EXISTS to_stock_id_orders_idx ON orders (to_stock_id);