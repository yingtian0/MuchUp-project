CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL REFERENCES users(id),

    sender_type TEXT NOT NULL DEFAULT 'USER',
    kind TEXT NOT NULL DEFAULT 'TEXT',
    status TEXT NOT NULL DEFAULT 'SENT',

    text TEXT,
    media_url TEXT,
    sticker_id TEXT,

    stream_id TEXT,
    sequence BIGINT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT messages_sender_type_check
        CHECK (sender_type IN ('USER', 'AI')),

    CONSTRAINT messages_kind_check
        CHECK (kind IN ('TEXT', 'MEDIA', 'STICKER')),

    CONSTRAINT messages_status_check
        CHECK (status IN ('PENDING', 'SENT', 'FAILED', 'DELETED')),

    CONSTRAINT messages_content_check
        CHECK (
            text IS NOT NULL
            OR media_url IS NOT NULL
            OR sticker_id IS NOT NULL
        )
);

ALTER TABLE room_members
    ADD CONSTRAINT room_members_last_read_message_id_fkey
    FOREIGN KEY (last_read_message_id) REFERENCES messages(id);

CREATE INDEX messages_room_created_at_idx
    ON messages (room_id, created_at, id);

CREATE INDEX messages_sender_created_at_idx
    ON messages (sender_id, created_at DESC);

CREATE UNIQUE INDEX messages_room_sequence_unique_idx
    ON messages (room_id, sequence)
    WHERE sequence IS NOT NULL;
