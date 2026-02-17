-- +migrate Up
ALTER TABLE users
ADD COLUMN IF NOT EXISTS email_verify BOOLEAN NOT NULL DEFAULT FALSE;

-- +migrate Down
ALTER TABLE users
DROP COLUMN IF EXISTS email_verify;