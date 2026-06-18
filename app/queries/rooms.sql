-- name: CreateRoom :one
INSERT INTO rooms (
    type,
    status,
    capacity,
    created_by
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetRoomByID :one
SELECT *
FROM rooms
WHERE id = $1;

-- name: ListRoomsByUserID :many
SELECT r.*
FROM rooms r
JOIN room_members rm ON rm.room_id = r.id
WHERE rm.user_id = $1
ORDER BY COALESCE(r.last_message_at, r.created_at) DESC;

-- name: UpdateRoomStatus :one
UPDATE rooms
SET
    status = $2,
    activated_at = CASE
        WHEN $2 = 'active' AND activated_at IS NULL THEN now()
        ELSE activated_at
    END,
    closed_at = CASE
        WHEN $2 = 'closed' AND closed_at IS NULL THEN now()
        ELSE closed_at
    END
WHERE id = $1
RETURNING *;

-- name: TouchRoomLastMessageAt :exec
UPDATE rooms
SET last_message_at = $2
WHERE id = $1;

-- name: FindWaitingRandomRoomWithAvailableSlots :one
SELECT r.*
FROM rooms r
WHERE r.type = 'random'
  AND r.status = 'waiting'
  AND (
      SELECT count(*)
      FROM room_members rm
      WHERE rm.room_id = r.id
        AND rm.status = 'JOINED'
  ) < r.capacity
ORDER BY r.created_at ASC
LIMIT 1;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;
