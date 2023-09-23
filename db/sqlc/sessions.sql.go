// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: sessions.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions(
    id,username,refresh_token,client_ip,expired_at,is_blocked
)
values ($1,$2,$3,$4,$5,$6)
RETURNING id, username, refresh_token, client_ip, is_blocked, expired_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	ClientIp     string    `json:"client_ip"`
	ExpiredAt    time.Time `json:"expired_at"`
	IsBlocked    bool      `json:"is_blocked"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.ClientIp,
		arg.ExpiredAt,
		arg.IsBlocked,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSessionById = `-- name: DeleteSessionById :one
DELETE FROM sessions WHERE id = $1
RETURNING id, username, refresh_token, client_ip, is_blocked, expired_at, created_at
`

func (q *Queries) DeleteSessionById(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, deleteSessionById, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSessionByUsername = `-- name: DeleteSessionByUsername :many
DELETE FROM sessions WHERE username = $1
RETURNING id, username, refresh_token, client_ip, is_blocked, expired_at, created_at
`

func (q *Queries) DeleteSessionByUsername(ctx context.Context, username string) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, deleteSessionByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.RefreshToken,
			&i.ClientIp,
			&i.IsBlocked,
			&i.ExpiredAt,
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

const getSessionFromId = `-- name: GetSessionFromId :one
SELECT id, username, refresh_token, client_ip, is_blocked, expired_at, created_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSessionFromId(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionFromId, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET 
    is_blocked = $1
WHERE   id = $2
RETURNING id, username, refresh_token, client_ip, is_blocked, expired_at, created_at
`

type UpdateSessionParams struct {
	IsBlocked bool      `json:"is_blocked"`
	ID        uuid.UUID `json:"id"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSession, arg.IsBlocked, arg.ID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}