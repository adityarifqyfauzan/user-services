
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id          bytea NOT NULL PRIMARY KEY,
    role_id     bytea NOT NULL,
    username    VARCHAR(50) UNIQUE NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    image       VARCHAR(255),
    name        VARCHAR(100),
    bio         VARCHAR(100),
    phone       VARCHAR(13),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(role_id)
        REFERENCES roles (id)
);
-- +migrate Down
DROP TABLE users;