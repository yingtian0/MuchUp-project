-- name: InsertUser :exec
INSERT INTO users (
    nickname,
    avatar_url,
    personality_profile,
    usage_purpose,
    status
) VALUES (
    $1, $2, $3, $4, $5
)
;

-- name: InsertUserWithID :exec
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
;

-- name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1
  AND deleted_at IS NULL;

-- name: FindAllUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
SET
    nickname = $2,
    avatar_url = $3,
    personality_profile = $4,
    usage_purpose = $5,
    status = $6,
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET
    status = 'DELETED',
    deleted_at = now(),
    updated_at = now()
WHERE id = $1
  AND deleted_at IS NULL;
