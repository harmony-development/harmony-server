// Code generated by sqlc. DO NOT EDIT.
// source: sessions.sql

package queries

import (
	"context"
)

const addSession = `-- name: AddSession :exec
INSERT INTO Sessions (User_ID,
                      Session)
VALUES ($1, $2)
`

type AddSessionParams struct {
	UserID  uint64 `json:"user_id"`
	Session string `json:"session"`
}

func (q *Queries) AddSession(ctx context.Context, arg AddSessionParams) error {
	_, err := q.exec(ctx, q.addSessionStmt, addSession, arg.UserID, arg.Session)
	return err
}

const sessionToUserID = `-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1
`

func (q *Queries) SessionToUserID(ctx context.Context, session string) (uint64, error) {
	row := q.queryRow(ctx, q.sessionToUserIDStmt, sessionToUserID, session)
	var user_id uint64
	err := row.Scan(&user_id)
	return user_id, err
}
