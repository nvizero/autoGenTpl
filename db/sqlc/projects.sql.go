// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: projects.sql

package db

import (
	"context"
	"database/sql"
)

const checkProject = `-- name: CheckProject :one
SELECT id, name, port, is_gen, created_at FROM projects
WHERE name like $1 or port = $2
`

type CheckProjectParams struct {
	Name sql.NullString `json:"name"`
	Port sql.NullInt32  `json:"port"`
}

func (q *Queries) CheckProject(ctx context.Context, arg CheckProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, checkProject, arg.Name, arg.Port)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Port,
		&i.IsGen,
		&i.CreatedAt,
	)
	return i, err
}

const createProjects = `-- name: CreateProjects :one
INSERT INTO projects (
    name,
    port,
    is_gen	
) VALUES (
  $1, $2, $3
) RETURNING id, name, port, is_gen, created_at
`

type CreateProjectsParams struct {
	Name  sql.NullString `json:"name"`
	Port  sql.NullInt32  `json:"port"`
	IsGen sql.NullInt32  `json:"is_gen"`
}

func (q *Queries) CreateProjects(ctx context.Context, arg CreateProjectsParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProjects, arg.Name, arg.Port, arg.IsGen)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Port,
		&i.IsGen,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, port, is_gen, created_at FROM projects
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Port,
		&i.IsGen,
		&i.CreatedAt,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, port, is_gen, created_at FROM projects
ORDER BY id desc
LIMIT $1
OFFSET $2
`

type ListProjectsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProjects(ctx context.Context, arg ListProjectsParams) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Port,
			&i.IsGen,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const truncateProject = `-- name: TruncateProject :exec
TRUNCATE TABLE projects
`

func (q *Queries) TruncateProject(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncateProject)
	return err
}
