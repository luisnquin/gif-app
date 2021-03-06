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

-- name: CreatePost :one
INSERT INTO posts(profile_id, external_source, description) VALUES($1, $2, $3) RETURNING *;

-- name: GetFullProfile :one
SELECT 
    u.id, u.username, u.firstname, u.lastname, u.email, u.roles, u.birthday, u.created_at, u.updated_at, p.last_connection
FROM users AS u INNER JOIN profiles AS p ON p.id = u.id WHERE u.id = $1; 

-- name: GetFullProfileByUsername :one
SELECT 
    u.id, u.username, u.firstname, u.lastname, u.email, u.roles, u.birthday, u.created_at, u.updated_at, p.last_connection
FROM users AS u INNER JOIN profiles AS p ON p.id = u.id WHERE u.username = $1; 

-- name: CreatePostWithTags :one
INSERT INTO posts(profile_id, external_source, description, tags) VALUES($1, $2, $3, $4) RETURNING *;

-- name: CreateMention :exec
INSERT INTO mentions(source, target) VALUES($1, $2);