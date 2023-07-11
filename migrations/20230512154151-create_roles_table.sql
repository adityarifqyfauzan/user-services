
-- +migrate Up
CREATE TABLE IF NOT EXISTS roles (
   id           bytea NOT NULL PRIMARY KEY,
   name         VARCHAR(50) NOT NULL,
   slug         VARCHAR(50) NOT NULL,
   created_at   TIMESTAMP NOT NULL DEFAULT now(),
   updated_at   TIMESTAMP NOT NULL DEFAULT now()
);
-- +migrate Down
DROP TABLE roles;