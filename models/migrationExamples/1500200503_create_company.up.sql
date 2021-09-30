CREATE TABLE IF NOT EXISTS companies (
    company_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_name VARCHAR (300),
    email VARCHAR (300),
    address VARCHAR (300),
    updated_at BIGINT
);