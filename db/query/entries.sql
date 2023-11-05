-- name: createEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY account_id
LIMIT $1
OFFSET $2;


-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;
