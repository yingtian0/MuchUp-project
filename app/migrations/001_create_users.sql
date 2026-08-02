CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  nickname VARCHAR(50) NOT NULL,
  avatar_url VARCHAR(500),
  personality_profile JSONB NOT NULL DEFAULT '{}'::jsonb,
  usage_purpose VARCHAR(255),

  status TEXT NOT NULL DEFAULT 'ACTIVE',

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ,

  CONSTRAINT users_status_check
      CHECK (status IN ('ACTIVE', 'SUSPENDED', 'DELETED'))
);
