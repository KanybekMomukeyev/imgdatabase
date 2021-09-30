CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_uuid VARCHAR (400) UNIQUE,
    user_type INTEGER,
    user_image_path VARCHAR (400),
    first_name VARCHAR (400),
    second_name VARCHAR (400),
    email VARCHAR (400) UNIQUE,
    password VARCHAR (400),
    phone_number VARCHAR (400),
    address VARCHAR (400),
    updated_at BIGINT
);