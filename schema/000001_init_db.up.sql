CREATE TABLE articles
(
    id uuid NOT NULL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    description VARCHAR(256),
    tags VARCHAR(100)[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);