-- name: CreateUser :one
INSERT INTO users (
    nickname,
    avatar_url,
    personality_profile,
    usage_purpose,
    status
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateUserWithID :one
INSERT INTO users (
    id,
    nickname,
    avatar_url,
    personality_profile,
    usage_purpose,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
  AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUserProfile :one
UPDATE users
SET
    nickname = $2,
    avatar_url = $3,
    personality_profile = $4,
    usage_purpose = $5,
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;

-- name: UpdateUserStatus :one
UPDATE users
SET
    status = $2,
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteUser :exec
UPDATE users
SET
    status = 'DELETED',
    deleted_at = now(),
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL;
