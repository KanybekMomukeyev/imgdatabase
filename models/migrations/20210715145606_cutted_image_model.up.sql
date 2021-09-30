CREATE TABLE IF NOT EXISTS cutted_images (
    image_id BIGSERIAL PRIMARY KEY NOT NULL,
    docmodel_id BIGINT,
    cutted_image_state INTEGER DEFAULT 1000,
    company_id BIGINT,
    folder_id BIGINT,
    parsed_image_path varchar (500),
    folder_name varchar (500),
    second_name varchar (500),
    phone_number varchar (500),
    address varchar (500),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_cuttedimages_idx ON cutted_images (company_id);
CREATE INDEX IF NOT EXISTS folder_id_cuttedimages_idx ON cutted_images (folder_id);
CREATE INDEX IF NOT EXISTS docmodel_id_cuttedimages_idx ON cutted_images (docmodel_id);