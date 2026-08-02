-- name: FindRoomMemberByRoomIDAndUserID :one
SELECT *
FROM room_members
WHERE room_id = $1
  AND user_id = $2;

-- name: FindAllRoomMembersByRoomID :many
SELECT *
FROM room_members
WHERE room_id = $1
ORDER BY joined_at ASC;

-- name: FindAllJoinedRoomMembersByRoomID :many
SELECT *
FROM room_members
WHERE room_id = $1
  AND status = 'JOINED'
ORDER BY joined_at ASC;

-- name: UpdateRoomMember :exec
UPDATE room_members
SET
    role = $3,
    status = $4,
    left_at = $5,
    last_read_message_id = $6
WHERE room_id = $1
  AND user_id = $2;
