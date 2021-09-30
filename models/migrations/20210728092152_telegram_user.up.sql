CREATE TABLE IF NOT EXISTS telegram_users (
    telegram_user_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT,
    company_id BIGINT,
    telegram_id BIGINT,
    first_name varchar (400),
    second_name varchar (400),
    phone_number varchar (400),
    tegram_account varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_tegram_users_idx ON telegram_users (company_id);
