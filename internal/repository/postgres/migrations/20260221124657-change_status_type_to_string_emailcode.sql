-- +migrate Up

ALTER TABLE email_codes
ADD COLUMN status_new text NOT NULL DEFAULT 'active';

UPDATE email_codes
SET status_new = CASE
  WHEN status = TRUE THEN 'verified'
  ELSE 'active'
END;

UPDATE email_codes
SET status_new = 'expired'
WHERE expiration_date < NOW()
  AND status_new = 'active';

ALTER TABLE email_codes DROP COLUMN status;
ALTER TABLE email_codes RENAME COLUMN status_new TO status;

ALTER TABLE email_codes ALTER COLUMN status SET DEFAULT 'active';

ALTER TABLE email_codes
ADD CONSTRAINT email_codes_status_check
CHECK (status IN ('active', 'verified', 'expired'));

-- (اختیاری ولی پیشنهادی)
CREATE INDEX IF NOT EXISTS idx_email_codes_email_status ON email_codes(email, status);


-- +migrate Down

ALTER TABLE email_codes
DROP CONSTRAINT IF EXISTS email_codes_status_check;

DROP INDEX IF EXISTS idx_email_codes_email_status;

ALTER TABLE email_codes
ADD COLUMN status_old boolean NOT NULL DEFAULT FALSE;

UPDATE email_codes
SET status_old = CASE
  WHEN status = 'verified' THEN TRUE
  ELSE FALSE
END;

ALTER TABLE email_codes DROP COLUMN status;
ALTER TABLE email_codes RENAME COLUMN status_old TO status;

-- ⚠️ اینجا: اگر قبلاً DEFAULT TRUE داشتی، همین بمونه
ALTER TABLE email_codes ALTER COLUMN status SET DEFAULT TRUE;