CREATE TABLE IF NOT EXISTS folders (
    folder_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    docmodel_id BIGINT,
    word_count BIGINT,
    parsed_image_count BIGINT,
    folder_image_path varchar (400),
    folder_name varchar (400),
    contact_fname varchar (400),
    phone_number varchar (400),
    address varchar (400),
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_folders_idx ON folders (company_id);
CREATE INDEX IF NOT EXISTS word_count_folders_idx ON folders (word_count);
CREATE INDEX IF NOT EXISTS docmodel_id_folders_idx ON folders (docmodel_id);