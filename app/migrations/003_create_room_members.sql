CREATE TABLE room_members (
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    role TEXT NOT NULL DEFAULT 'MEMBER',
    status TEXT NOT NULL DEFAULT 'JOINED',

    joined_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    left_at TIMESTAMPTZ,
    last_read_message_id UUID,

    PRIMARY KEY (room_id, user_id),

    CONSTRAINT room_members_role_check
        CHECK (role IN ('OWNER', 'MEMBER', 'AI')),

    CONSTRAINT room_members_status_check
        CHECK (status IN ('JOINED', 'LEFT', 'KICKED'))
);

CREATE INDEX room_members_user_id_idx
    ON room_members (user_id);

CREATE INDEX room_members_room_status_idx
    ON room_members (room_id, status);
