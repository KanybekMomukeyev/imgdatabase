CREATE TABLE IF NOT EXISTS customers (
    customer_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT,
    company_id BIGINT,
    stock_id BIGINT,
    customer_image_path varchar (400),
    first_name varchar (400),
    second_name varchar (400),
    phone_number varchar (400),
    address varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_customers_idx ON customers (company_id);
CREATE INDEX IF NOT EXISTS stock_id_customers_idx ON customers (stock_id);