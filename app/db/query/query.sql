-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY created_at DESC;

-- name: CreateEntry :one
INSERT INTO entries (title, message, mood) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateEntry :one
UPDATE entries SET title = $1, message = $2 , mood = $3 WHERE id = $4 RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entries WHERE id = $1 RETURNING id;