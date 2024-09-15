CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email citext UNIQUE NOT NULL,
  password_hash bytea NOT NULL,
  activated BOOL NOT NULL,
  created_at timestamp(0) WITH time ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) WITH time ZONE NOT NULL DEFAULT NOW()
);