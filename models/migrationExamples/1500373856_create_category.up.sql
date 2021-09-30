CREATE TABLE IF NOT EXISTS categories (
    category_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    category_name varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_categories_idx ON categories (company_id);


