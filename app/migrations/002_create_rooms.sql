CREATE TABLE rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    type TEXT NOT NULL DEFAULT 'random',
    status TEXT NOT NULL DEFAULT 'waiting',
    capacity INTEGER NOT NULL DEFAULT 5,

    created_by UUID REFERENCES users(id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    activated_at TIMESTAMPTZ,
    closed_at TIMESTAMPTZ,
    last_message_at TIMESTAMPTZ,
    last_ai_intervened_at TIMESTAMPTZ,

    CONSTRAINT rooms_type_check
        CHECK (type IN ('random', 'fixed')),

    CONSTRAINT rooms_status_check
        CHECK (status IN ('waiting', 'active', 'closed')),

    CONSTRAINT rooms_capacity_check
        CHECK (capacity > 0)
);

CREATE INDEX rooms_matching_lookup_idx
    ON rooms (type, status, created_at);
