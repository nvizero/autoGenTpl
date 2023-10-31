-- name: CreateProjects :one
INSERT INTO projects (
    name,
    port,
    is_gen	
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProjectByName :one
SELECT * FROM projects
WHERE name like $1 ;

-- name: GetProject :one
SELECT * FROM projects
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT * FROM projects
ORDER BY id desc
LIMIT $1
OFFSET $2;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1;

-- name: TruncateProject :exec 
TRUNCATE TABLE projects;
