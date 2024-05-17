-- name: GetClient :one
SELECT * FROM client
WHERE client_id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM client
ORDER BY given_name;

-- name: GetClientByEmail :one
SELECT email, password, given_name, surname
FROM client
WHERE email = $1;

-- name: CreateClient :one
INSERT INTO client (email, password, given_name, surname, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateClientById :exec
UPDATE client
SET email = $2, password = $3, given_name = $4, surname = $5, updated_at = $6
WHERE client_id = $1;

-- name: DeleteClient :exec
DELETE FROM client
WHERE client_id = $1;

-- name: DeleteAllClients :execresult
DELETE FROM client;
