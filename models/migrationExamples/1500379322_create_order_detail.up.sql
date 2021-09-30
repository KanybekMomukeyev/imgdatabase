CREATE TABLE IF NOT EXISTS orderdetails (
    order_detail_id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    company_id BIGINT,
    stock_id BIGINT,
    order_detail_date BIGINT,
    is_last INTEGER,
    billing_no varchar (400),
    orderdetail_comment varchar (400),
    product_id BIGINT,
    price REAL,
	order_quantity REAL,
	product_quantity REAL,
	uuid varchar (400),
	updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_order_details_idx ON orderdetails (company_id);
--CREATE INDEX IF NOT EXISTS stock_id_order_details_idx ON orderdetails (stock_id);
CREATE INDEX IF NOT EXISTS order_id_order_details_idx ON orderdetails (order_id);
--CREATE INDEX IF NOT EXISTS product_id_order_details_idx ON orderdetails (product_id);
--CREATE INDEX IF NOT EXISTS order_detail_date_order_details_idx ON orderdetails (order_detail_date);