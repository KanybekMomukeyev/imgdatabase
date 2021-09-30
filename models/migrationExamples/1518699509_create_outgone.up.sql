CREATE TABLE IF NOT EXISTS outgones (
    outgone_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    outgone_name varchar (500),
    updated_at BIGINT
);