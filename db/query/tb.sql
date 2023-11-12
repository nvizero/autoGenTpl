-- name: CreateTb :one
INSERT INTO tb (
    name,
    project_id,
    describe
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTb :one
SELECT * FROM tb 
WHERE id = $1 LIMIT 1;

-- name: ListTb :many
SELECT * FROM tb 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CheckTb :one
SELECT id,name FROM tb 
WHERE project_id = $1
AND name = $2;

-- name: WhereTbByPID :many
SELECT * FROM tb 
WHERE project_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTb :one
UPDATE tb 
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTb :exec
DELETE FROM tb 
WHERE id = $1;

-- name: DeleteTbByPID :exec
DELETE FROM tb 
WHERE project_id = $1;

-- name: TruncateTB :exec 
TRUNCATE TABLE tb;

