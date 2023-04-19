// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: chat_snapshot.sql

package sqlc_queries

import (
	"context"
	"encoding/json"
	"time"
)

const chatSnapshotByID = `-- name: ChatSnapshotByID :one
SELECT id, uuid, user_id, title, summary, model, tags, conversation, created_at FROM chat_snapshot WHERE id = $1
`

func (q *Queries) ChatSnapshotByID(ctx context.Context, id int32) (ChatSnapshot, error) {
	row := q.db.QueryRowContext(ctx, chatSnapshotByID, id)
	var i ChatSnapshot
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Model,
		&i.Tags,
		&i.Conversation,
		&i.CreatedAt,
	)
	return i, err
}

const chatSnapshotByUUID = `-- name: ChatSnapshotByUUID :one
SELECT id, uuid, user_id, title, summary, model, tags, conversation, created_at FROM chat_snapshot WHERE uuid = $1
`

func (q *Queries) ChatSnapshotByUUID(ctx context.Context, uuid string) (ChatSnapshot, error) {
	row := q.db.QueryRowContext(ctx, chatSnapshotByUUID, uuid)
	var i ChatSnapshot
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Model,
		&i.Tags,
		&i.Conversation,
		&i.CreatedAt,
	)
	return i, err
}

const chatSnapshotMetaByUserID = `-- name: ChatSnapshotMetaByUserID :many
SELECT uuid, title, summary, tags, created_at
FROM chat_snapshot WHERE user_id = $1
order by id desc
`

type ChatSnapshotMetaByUserIDRow struct {
	Uuid      string
	Title     string
	Summary   string
	Tags      json.RawMessage
	CreatedAt time.Time
}

func (q *Queries) ChatSnapshotMetaByUserID(ctx context.Context, userID int32) ([]ChatSnapshotMetaByUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, chatSnapshotMetaByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatSnapshotMetaByUserIDRow
	for rows.Next() {
		var i ChatSnapshotMetaByUserIDRow
		if err := rows.Scan(
			&i.Uuid,
			&i.Title,
			&i.Summary,
			&i.Tags,
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

const createChatSnapshot = `-- name: CreateChatSnapshot :one
INSERT INTO chat_snapshot (uuid, user_id, title, model, summary, tags, conversation )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, uuid, user_id, title, summary, model, tags, conversation, created_at
`

type CreateChatSnapshotParams struct {
	Uuid         string
	UserID       int32
	Title        string
	Model        string
	Summary      string
	Tags         json.RawMessage
	Conversation json.RawMessage
}

func (q *Queries) CreateChatSnapshot(ctx context.Context, arg CreateChatSnapshotParams) (ChatSnapshot, error) {
	row := q.db.QueryRowContext(ctx, createChatSnapshot,
		arg.Uuid,
		arg.UserID,
		arg.Title,
		arg.Model,
		arg.Summary,
		arg.Tags,
		arg.Conversation,
	)
	var i ChatSnapshot
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Model,
		&i.Tags,
		&i.Conversation,
		&i.CreatedAt,
	)
	return i, err
}

const deleteChatSnapshot = `-- name: DeleteChatSnapshot :exec
DELETE FROM chat_snapshot WHERE id = $1
`

func (q *Queries) DeleteChatSnapshot(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteChatSnapshot, id)
	return err
}

const listChatSnapshots = `-- name: ListChatSnapshots :many
SELECT id, uuid, user_id, title, summary, model, tags, conversation, created_at FROM chat_snapshot ORDER BY id
`

func (q *Queries) ListChatSnapshots(ctx context.Context) ([]ChatSnapshot, error) {
	rows, err := q.db.QueryContext(ctx, listChatSnapshots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatSnapshot
	for rows.Next() {
		var i ChatSnapshot
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.UserID,
			&i.Title,
			&i.Summary,
			&i.Model,
			&i.Tags,
			&i.Conversation,
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

const updateChatSnapshot = `-- name: UpdateChatSnapshot :one
UPDATE chat_snapshot
SET uuid = $2, user_id = $3, title = $4, summary = $5, tags = $6, conversation = $7, created_at = $8
WHERE id = $1
RETURNING id, uuid, user_id, title, summary, model, tags, conversation, created_at
`

type UpdateChatSnapshotParams struct {
	ID           int32
	Uuid         string
	UserID       int32
	Title        string
	Summary      string
	Tags         json.RawMessage
	Conversation json.RawMessage
	CreatedAt    time.Time
}

func (q *Queries) UpdateChatSnapshot(ctx context.Context, arg UpdateChatSnapshotParams) (ChatSnapshot, error) {
	row := q.db.QueryRowContext(ctx, updateChatSnapshot,
		arg.ID,
		arg.Uuid,
		arg.UserID,
		arg.Title,
		arg.Summary,
		arg.Tags,
		arg.Conversation,
		arg.CreatedAt,
	)
	var i ChatSnapshot
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Model,
		&i.Tags,
		&i.Conversation,
		&i.CreatedAt,
	)
	return i, err
}

const updateChatSnapshotMetaByUUID = `-- name: UpdateChatSnapshotMetaByUUID :exec
UPDATE chat_snapshot
SET title = $2, summary = $3
WHERE uuid = $1
`

type UpdateChatSnapshotMetaByUUIDParams struct {
	Uuid    string
	Title   string
	Summary string
}

func (q *Queries) UpdateChatSnapshotMetaByUUID(ctx context.Context, arg UpdateChatSnapshotMetaByUUIDParams) error {
	_, err := q.db.ExecContext(ctx, updateChatSnapshotMetaByUUID, arg.Uuid, arg.Title, arg.Summary)
	return err
}