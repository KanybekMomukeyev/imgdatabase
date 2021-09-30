ALTER TABLE cutted_image_translates ADD COLUMN IF NOT EXISTS tsv tsvector;
UPDATE cutted_image_translates SET tsv = setweight(to_tsvector(translated_word), 'A');
CREATE INDEX cutted_image_translates_tsv ON cutted_image_translates USING GIN(tsv);