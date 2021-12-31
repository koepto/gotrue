-- adds parent column

ALTER TABLE auth.refresh_token
ADD COLUMN IF NOT EXISTS parent varchar(255) NULL,
ADD CONSTRAINT refresh_token_token_unique UNIQUE ("token"),
ADD CONSTRAINT refresh_token_parent_fkey FOREIGN KEY (parent) REFERENCES auth.refresh_token("token");

CREATE INDEX IF NOT EXISTS refresh_token_parent_idx ON refresh_token USING btree (parent);
