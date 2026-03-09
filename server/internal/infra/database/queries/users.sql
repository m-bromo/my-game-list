-- name: GetByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: Save :exec
insert into users (id, username, email, password)
values ($1, $2, $3, $4);