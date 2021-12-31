-- adds identity table 

CREATE TABLE IF NOT EXISTS auth.identity (
    id text NOT NULL,
    user_id uuid NOT NULL,
    identity_data JSONB NOT NULL,
    provider text NOT NULL,
    last_sign_in_at timestamptz NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    CONSTRAINT identity_pkey PRIMARY KEY (provider, id),
    CONSTRAINT identity_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth.user(id) ON DELETE CASCADE
);
COMMENT ON TABLE auth.identity is 'Auth: Stores identity associated to a user.';
