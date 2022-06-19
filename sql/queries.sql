-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, firstname, lastname, email, password, roles) 
VALUES($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UserExistsByUsernameOrEmail :one
SELECT exists(SELECT * FROM users WHERE username=$1 OR email=$2);

-- name: ChangePasswordByID :exec
UPDATE users SET password = $1 WHERE id = $2;

-- name: UserExists :one
SELECT exists(SELECT * FROM users WHERE id = $1);

-- name: GetUserByUsernameOrEmail :one
SELECT * FROM users WHERE username = $1 OR email = $2;

-- name: CreateProfile :one
INSERT INTO profiles(id) VALUES($1) RETURNING *;
