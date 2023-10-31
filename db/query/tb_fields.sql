-- name: CreateTbField :one
INSERT INTO tb_fields (
    table_id,
    field_name,
    show_name,
    migration,
    model_type,
    is_require
) VALUES (
  $1, $2, $3, $4, $5, $6
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
