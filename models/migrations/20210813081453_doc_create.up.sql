CREATE TABLE IF NOT EXISTS docmodels (
    docmodel_id BIGSERIAL PRIMARY KEY NOT NULL,
    company_id BIGINT,
    user_id BIGINT,
    parsed_image_count INTEGER,
    stat_count INTEGER,
    docmodel_name varchar (400),
    summary TEXT,
    comments TEXT,
    descriptionn TEXT,
    updated_at BIGINT
);

CREATE INDEX IF NOT EXISTS company_id_docmodels_idx ON docmodels (company_id);
CREATE INDEX IF NOT EXISTS user_id_docmodels_idx ON docmodels (user_id);
