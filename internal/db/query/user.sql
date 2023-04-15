-- name: CreateUser :one
INSERT INTO "users" (
	firstname,
	lastname,
	username,
	email,
	password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT
	id,
	firstname,
	lastname,
	username,
	email,
	password,
	created_at,
	updated_at
FROM "users" WHERE "id" = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT
	id,
	firstname,
	lastname,
	username,
	email,
	password,
	created_at,
	updated_at
FROM "users" WHERE "email" = $1;

-- name: GetAllUsers :many
SELECT
	id,
	firstname,
	lastname,
	username,
	email,
	password,
	created_at,
	updated_at
FROM "users";

-- name: UpdateUser :one
UPDATE "users"
SET
	firstname = $2,
	lastname = $3,
	username = $4,
	email = $5,
	password = $6,
	updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM "users"
WHERE id = $1
RETURNING id;