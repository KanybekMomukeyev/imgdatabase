CREATE TABLE IF NOT EXISTS cutted_image_translates (
    cutted_image_translate_id BIGSERIAL PRIMARY KEY NOT NULL,
    cutted_image_id BIGINT,
    telegram_user_id BIGINT,
    translated_word varchar (400),
    comments varchar (400),
    summary varchar (400),
    accept_status INTEGER,
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS cutted_image_id_image_translates_idx ON cutted_image_translates (cutted_image_id);
CREATE INDEX IF NOT EXISTS telegram_user_id_image_translates_idx ON cutted_image_translates (telegram_user_id);