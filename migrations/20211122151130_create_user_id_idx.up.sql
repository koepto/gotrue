-- create index on identity.user_id

CREATE INDEX IF NOT EXISTS identity_user_id_idx ON identity using btree (user_id);
