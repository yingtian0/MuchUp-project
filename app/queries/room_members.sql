-- name: GetRoomMember :one
SELECT *
FROM room_members
WHERE room_id = $1
  AND user_id = $2;

-- name: ListRoomMembers :many
SELECT *
FROM room_members
WHERE room_id = $1
ORDER BY joined_at ASC;

-- name: ListJoinedRoomMembers :many
SELECT *
FROM room_members
WHERE room_id = $1
  AND status = 'JOINED'
ORDER BY joined_at ASC;

-- name: LeaveRoom :one
UPDATE room_members
SET
    status = 'LEFT',
    left_at = now()
WHERE room_id = $1
  AND user_id = $2
RETURNING *;

-- name: UpdateLastReadMessage :one
UPDATE room_members
SET last_read_message_id = $3
WHERE room_id = $1
  AND user_id = $2
RETURNING *;
