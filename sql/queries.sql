-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, firstname, lastname, email, password, role) 
VALUES($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UserExists :one
SELECT exists(SELECT * FROM users WHERE username=$1 OR email=$2);

-- name: GetUserByUsernameOrEmail :one
SELECT * FROM users WHERE username = $1 OR email = $2;