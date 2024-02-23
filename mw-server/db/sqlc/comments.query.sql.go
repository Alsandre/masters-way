// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: comments.query.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments(
    created_at,
    updated_at,
    description,
    owner_uuid,
    day_report_uuid
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING uuid, created_at, updated_at, description, owner_uuid, day_report_uuid
`

type CreateCommentParams struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Description   string    `json:"description"`
	OwnerUuid     uuid.UUID `json:"owner_uuid"`
	DayReportUuid uuid.UUID `json:"day_report_uuid"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.queryRow(ctx, q.createCommentStmt, createComment,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Description,
		arg.OwnerUuid,
		arg.DayReportUuid,
	)
	var i Comment
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.OwnerUuid,
		&i.DayReportUuid,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE uuid = $1
`

func (q *Queries) DeleteComment(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteCommentStmt, deleteComment, argUuid)
	return err
}

const getListCommentsByDayReportId = `-- name: GetListCommentsByDayReportId :many
SELECT
    comments.uuid,
    comments.created_at,
    comments.updated_at,
    comments.description,
    users.name AS owner_name,
    users.uuid AS owner_uuid
FROM comments
JOIN users ON comments.owner_uuid = users.uuid
WHERE day_report_uuid = $1
ORDER BY comments.created_at
`

type GetListCommentsByDayReportIdRow struct {
	Uuid        uuid.UUID `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	OwnerName   string    `json:"owner_name"`
	OwnerUuid   uuid.UUID `json:"owner_uuid"`
}

func (q *Queries) GetListCommentsByDayReportId(ctx context.Context, dayReportUuid uuid.UUID) ([]GetListCommentsByDayReportIdRow, error) {
	rows, err := q.query(ctx, q.getListCommentsByDayReportIdStmt, getListCommentsByDayReportId, dayReportUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListCommentsByDayReportIdRow{}
	for rows.Next() {
		var i GetListCommentsByDayReportIdRow
		if err := rows.Scan(
			&i.Uuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.OwnerName,
			&i.OwnerUuid,
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

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET
updated_at = coalesce($1, updated_at),
description = coalesce($2, description)
WHERE uuid = $3
RETURNING uuid, created_at, updated_at, description, owner_uuid, day_report_uuid
`

type UpdateCommentParams struct {
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	Description sql.NullString `json:"description"`
	Uuid        uuid.UUID      `json:"uuid"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.queryRow(ctx, q.updateCommentStmt, updateComment, arg.UpdatedAt, arg.Description, arg.Uuid)
	var i Comment
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.OwnerUuid,
		&i.DayReportUuid,
	)
	return i, err
}