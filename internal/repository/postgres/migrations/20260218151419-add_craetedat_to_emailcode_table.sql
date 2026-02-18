-- +migrate Up
ALTER TABLE email_codes
ADD COLUMN IF NOT EXISTS created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

CREATE INDEX IF NOT EXISTS idx_email_codes_created_at ON email_codes(created_at);

-- +migrate Down
DROP INDEX IF EXISTS idx_email_codes_created_at;

ALTER TABLE email_codes
DROP COLUMN IF EXISTS created_at;
