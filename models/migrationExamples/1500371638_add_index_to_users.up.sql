CREATE INDEX IF NOT EXISTS users_company_id_index ON users (company_id);
CREATE UNIQUE INDEX IF NOT EXISTS users_user_uuid_index ON users (user_uuid);
CREATE INDEX IF NOT EXISTS users_stock_id_index ON users (stock_id);