ALTER TABLE cutted_images ADD COLUMN IF NOT EXISTS cutted_image_type INTEGER DEFAULT 101010;
CREATE INDEX IF NOT EXISTS type_cutted_image_type_idx ON cutted_images (cutted_image_type);
CREATE INDEX IF NOT EXISTS state_cutted_image_type_idx ON cutted_images (cutted_image_state);