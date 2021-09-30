CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_uuid VARCHAR (300) UNIQUE,
    user_type INTEGER,
    user_image_path VARCHAR (300),
    first_name VARCHAR (300),
    second_name VARCHAR (300),
    email VARCHAR (300) UNIQUE,
    password VARCHAR (500),
    phone_number VARCHAR (300),
    address VARCHAR (300),
    updated_at BIGINT
);
