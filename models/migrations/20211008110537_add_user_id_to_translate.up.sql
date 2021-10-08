ALTER TABLE cutted_image_translates ADD COLUMN IF NOT EXISTS user_id BIGINT DEFAULT 0;
CREATE INDEX IF NOT EXISTS user_id_cutted_image_translates_idx ON cutted_image_translates (user_id);