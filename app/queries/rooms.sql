-- name: InsertRoom :one
INSERT INTO rooms (
    type,
    status,
    capacity,
    created_by
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: FindRoomByID :one
SELECT *
FROM rooms
WHERE id = $1;

-- name: FindAllRoomsByUserID :many
SELECT r.*
FROM rooms r
JOIN room_members rm ON rm.room_id = r.id
WHERE rm.user_id = $1
ORDER BY COALESCE(r.last_message_at, r.created_at) DESC;

-- name: UpdateRoom :one
UPDATE rooms
SET
    status = $2,
    type = $3,
    capacity = $4,
    activated_at = CASE
        WHEN $2 = 'active' AND activated_at IS NULL THEN now()
        ELSE $5
    END,
    closed_at = CASE
        WHEN $2 = 'closed' AND closed_at IS NULL THEN now()
        ELSE $6
    END,
    last_message_at = $7,
    last_ai_intervened_at = $8
WHERE id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;
