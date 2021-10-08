DROP INDEX IF EXISTS user_id_cutted_image_translates_idx;
ALTER TABLE cutted_image_translates DROP COLUMN IF EXISTS user_id;