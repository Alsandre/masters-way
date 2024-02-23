// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: from_user_mentoring_requests.query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFromUserMentoringRequest = `-- name: CreateFromUserMentoringRequest :one
INSERT INTO from_user_mentoring_requests(
    user_uuid,
    way_uuid
) VALUES (
    $1, $2
) RETURNING user_uuid, way_uuid
`

type CreateFromUserMentoringRequestParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	WayUuid  uuid.UUID `json:"way_uuid"`
}

func (q *Queries) CreateFromUserMentoringRequest(ctx context.Context, arg CreateFromUserMentoringRequestParams) (FromUserMentoringRequest, error) {
	row := q.queryRow(ctx, q.createFromUserMentoringRequestStmt, createFromUserMentoringRequest, arg.UserUuid, arg.WayUuid)
	var i FromUserMentoringRequest
	err := row.Scan(&i.UserUuid, &i.WayUuid)
	return i, err
}

const deleteFromUserMentoringRequestByIds = `-- name: DeleteFromUserMentoringRequestByIds :exec
DELETE FROM from_user_mentoring_requests
WHERE user_uuid = $1 AND way_uuid = $2
`

type DeleteFromUserMentoringRequestByIdsParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	WayUuid  uuid.UUID `json:"way_uuid"`
}

func (q *Queries) DeleteFromUserMentoringRequestByIds(ctx context.Context, arg DeleteFromUserMentoringRequestByIdsParams) error {
	_, err := q.exec(ctx, q.deleteFromUserMentoringRequestByIdsStmt, deleteFromUserMentoringRequestByIds, arg.UserUuid, arg.WayUuid)
	return err
}