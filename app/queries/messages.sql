-- name: InsertMessage :one
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

-- name: InsertTextMessage :one
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

-- name: FindMessageByID :one
SELECT *
FROM messages
WHERE id = $1
  AND deleted_at IS NULL;

-- name: FindAllMessagesByRoom :many
SELECT *
FROM messages
WHERE room_id = $1
  AND deleted_at IS NULL
ORDER BY created_at ASC, id ASC
LIMIT $2 OFFSET $3;

-- name: FindAllMessagesByUserID :many
SELECT *
FROM messages
WHERE sender_id = $1
  AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateMessage :one
UPDATE messages
SET
    status = $2,
    text = $3,
    media_url = $4,
    sticker_id = $5,
    stream_id = $6,
    sequence = $7,
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;

-- name: DeleteMessage :one
UPDATE messages
SET
    status = 'DELETED',
    deleted_at = now(),
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;
