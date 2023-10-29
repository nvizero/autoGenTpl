-- name: CreateTbField :one
INSERT INTO tb_fields (
    table_id,
    field_name,
    laravel_map
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteTbFieldByTableID :exec
DELETE FROM tb_fields
WHERE table_id = $1;

-- name: GetTFBytID :many
SELECT * FROM tb_fields 
WHERE table_id = $1;

-- name: GetTFByfID :many
SELECT * FROM tb_fields 
WHERE table_id = $1 ;

-- name: TruncateTbField :exec 
TRUNCATE TABLE tb_fields;

-- name: ListTbField :many
SELECT * FROM tb_fields
ORDER BY id desc
LIMIT $1
OFFSET $2;

