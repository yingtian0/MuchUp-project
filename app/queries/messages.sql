-- name: CreateMessage :one
INSERT INTO messages (
    room_id,
    sender_id,
    sender_type,
    kind,
    status,
    text,
    media_url,
    sticker_id,
    stream_id,
    sequence
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: CreateTextMessage :one
INSERT INTO messages (
    room_id,
    sender_id,
    sender_type,
    kind,
    status,
    text,
    stream_id,
    sequence
) VALUES (
    $1, $2, $3, 'TEXT', 'SENT', $4, $5, $6
)
RETURNING *;

-- name: GetMessageByID :one
SELECT *
FROM messages
WHERE id = $1
  AND deleted_at IS NULL;

-- name: ListMessagesByRoom :many
SELECT *
FROM messages
WHERE room_id = $1
  AND deleted_at IS NULL
ORDER BY created_at ASC, id ASC
LIMIT $2 OFFSET $3;

-- name: ListMessagesByUserID :many
SELECT *
FROM messages
WHERE sender_id = $1
  AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateMessageStatus :one
UPDATE messages
SET
    status = $2,
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteMessage :one
UPDATE messages
SET
    status = 'DELETED',
    deleted_at = now(),
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;
