-- adds confirmed at

ALTER TABLE auth.user
ADD COLUMN IF NOT EXISTS confirmed_at timestamptz GENERATED ALWAYS AS (LEAST ("user".email_confirmed_at, "user".phone_confirmed_at)) STORED;
