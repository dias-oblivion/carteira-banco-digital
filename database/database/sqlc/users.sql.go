// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package database

import (
	"context"
)

const getUsers = `-- name: GetUsers :many
SELECT id, name, document, email, password, role FROM users
ORDER BY id
LIMIT $1 OFFSET $2
`

type GetUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Document,
			&i.Email,
			&i.Password,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
