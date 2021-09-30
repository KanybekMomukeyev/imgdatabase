CREATE TABLE IF NOT EXISTS stockdetails (
    stock_detail_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    stock_id BIGINT,
    product_id BIGINT,
    sale_unit_price REAL,
    income_unit_price REAL,
    units_in_stock REAL,
    stock_detail_uuid VARCHAR (300),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_stockdetails_idx ON stockdetails (company_id);
CREATE INDEX IF NOT EXISTS stock_id_stockdetails_idx ON stockdetails (stock_id);
CREATE INDEX IF NOT EXISTS product_id_stockdetails_idx ON stockdetails (product_id);


