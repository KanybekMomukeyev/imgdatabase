CREATE TABLE IF NOT EXISTS products (
    product_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    product_type INTEGER,
    product_image_path varchar (400),
    product_name varchar (400),
    supplier_id BIGINT,
    category_id BIGINT,
    barcode VARCHAR (300),
    quantity_per_unit VARCHAR (300),
    sale_unit_price REAL,
    income_unit_price REAL,
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_products_idx ON products (company_id);
--CREATE INDEX IF NOT EXISTS supplier_id_products_idx ON products (supplier_id);
--CREATE INDEX IF NOT EXISTS category_id_products_idx ON products (category_id);
