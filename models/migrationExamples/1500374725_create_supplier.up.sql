CREATE TABLE IF NOT EXISTS suppliers (
    supplier_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    stock_id BIGINT,
    supplier_image_path varchar (400),
    company_name varchar (400),
    contact_fname varchar (400),
    phone_number varchar (400),
    address varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_suppliers_idx ON suppliers (company_id);
CREATE INDEX IF NOT EXISTS stock_id_suppliers_idx ON suppliers (stock_id);