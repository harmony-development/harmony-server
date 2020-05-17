// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package queries

import (
	"context"
	"database/sql"
)

const addUser = `-- name: AddUser :exec
INSERT INTO Users (User_ID, Email, Username, Avatar, Password)
VALUES ($1, $2, $3, $4, $5)
`

type AddUserParams struct {
	UserID   uint64         `json:"user_id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
	Password []byte         `json:"password"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.exec(ctx, q.addUserStmt, addUser,
		arg.UserID,
		arg.Email,
		arg.Username,
		arg.Avatar,
		arg.Password,
	)
	return err
}

const emailExists = `-- name: EmailExists :one
SELECT Email
FROM Users
WHERE Email = $1
`

func (q *Queries) EmailExists(ctx context.Context, email string) (string, error) {
	row := q.queryRow(ctx, q.emailExistsStmt, emailExists, email)
	var email string
	err := row.Scan(&email)
	return email, err
}

const getUser = `-- name: GetUser :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE Email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.Avatar,
		&i.Password,
	)
	return i, err
}
