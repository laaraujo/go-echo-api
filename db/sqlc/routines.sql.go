// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: routines.sql

package db

import (
	"context"
)

const createRoutine = `-- name: CreateRoutine :one
INSERT INTO routines (
  name
) VALUES (
  $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreateRoutine(ctx context.Context, name string) (Routine, error) {
	row := q.db.QueryRow(ctx, createRoutine, name)
	var i Routine
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteRoutine = `-- name: DeleteRoutine :exec
DELETE FROM routines
WHERE id = $1
`

func (q *Queries) DeleteRoutine(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRoutine, id)
	return err
}

const getRoutine = `-- name: GetRoutine :one
SELECT id, name, created_at FROM routines
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoutine(ctx context.Context, id int32) (Routine, error) {
	row := q.db.QueryRow(ctx, getRoutine, id)
	var i Routine
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listRoutines = `-- name: ListRoutines :many
SELECT id, name, created_at FROM routines
ORDER BY name
`

func (q *Queries) ListRoutines(ctx context.Context) ([]Routine, error) {
	rows, err := q.db.Query(ctx, listRoutines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Routine{}
	for rows.Next() {
		var i Routine
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoutine = `-- name: UpdateRoutine :exec
UPDATE routines
  set name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateRoutineParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateRoutine(ctx context.Context, arg UpdateRoutineParams) error {
	_, err := q.db.Exec(ctx, updateRoutine, arg.ID, arg.Name)
	return err
}
