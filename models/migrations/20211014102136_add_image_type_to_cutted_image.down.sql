ALTER TABLE cutted_images DROP COLUMN IF EXISTS cutted_image_type;
DROP INDEX IF EXISTS type_cutted_image_type_idx;
DROP INDEX IF EXISTS state_cutted_image_type_idx;
