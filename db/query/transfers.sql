-- name: GetTranfersfrom :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id;

-- name: GetTranfersto :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id;

-- name: ListAlltransfer :many
SELECT * FROM transfers
ORDER BY id;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
   to_account_id,
   amount
) VALUES (
  $1, $2, $3
)RETURNING *;

-- name: UpdateTransfer :exec
UPDATE transfers
set completed = $2
WHERE id = $1
RETURNING *;