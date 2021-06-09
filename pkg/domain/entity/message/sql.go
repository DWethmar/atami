// This file was generated by robots; DO NOT EDIT.
// run: 'make generate' to regenerate this file.

package message

import (
	"database/sql"
	"github.com/dwethmar/atami/pkg/database"

	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

// Row needs to be implemented in the the map function.
type Row interface {
	Scan(dest ...interface{}) error
	Err() error
}

// selectMessages sql query
var selectMessagesSQL = `SELECT
	message.id,
	message.uid,
	message.text,
	message.created_by_user_id,
	message.created_at,
	message.updated_at,
	app_user.id,
	app_user.uid,
	app_user.username
FROM message
LEFT JOIN app_user ON message.created_by_user_id = app_user.id
ORDER BY message.created_at DESC
LIMIT $1
OFFSET $2`

func querySelectMessages(
	db database.Transaction,
	limit uint,
	offset uint,
) ([]*Message, error) {
	rows, err := db.Query(
		selectMessagesSQL,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]*Message, 0)
	for rows.Next() {
		if entry, err := messageWithUserRowMap(rows); err == nil {
			entries = append(entries, entry)
		} else {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

// selectMessageByID sql query
var selectMessageByIDSQL = `SELECT
	message.id,
	message.uid,
	message.text,
	message.created_by_user_id,
	message.created_at,
	message.updated_at,
	app_user.id,
	app_user.uid,
	app_user.username
FROM message
LEFT JOIN app_user ON message.created_by_user_id = app_user.id
WHERE message.id = $1`

func queryRowSelectMessageByID(
	db database.Transaction,
	ID entity.ID,
) (*Message, error) {
	return messageWithUserRowMap(db.QueryRow(
		selectMessageByIDSQL,
		ID,
	))
}

// selectMessageByUID sql query
var selectMessageByUIDSQL = `SELECT
	message.id,
	message.uid,
	message.text,
	message.created_by_user_id,
	message.created_at,
	message.updated_at,
	app_user.id,
	app_user.uid,
	app_user.username
FROM message
LEFT JOIN app_user ON message.created_by_user_id = app_user.id
WHERE message.uid = $1`

func queryRowSelectMessageByUID(
	db database.Transaction,
	UID string,
) (*Message, error) {
	return messageWithUserRowMap(db.QueryRow(
		selectMessageByUIDSQL,
		UID,
	))
}

// deleteMessage sql query
var deleteMessageSQL = `DELETE FROM message
WHERE message.id = $1`

func execDeleteMessage(
	db database.Transaction,
	ID entity.ID,
) (sql.Result, error) {
	return db.Exec(
		deleteMessageSQL,
		ID,
	)
}

// insertMessage sql query
var insertMessageSQL = `INSERT INTO message
(
	uid,
	text,
	created_by_user_id,
	created_at,
	updated_at
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
)
RETURNING id`

func queryRowInsertMessage(
	db database.Transaction,
	UID string,
	text string,
	CreatedByUserID int,
	createdAt time.Time,
	updatedAt time.Time,
) (entity.ID, error) {
	return insertRowMap(db.QueryRow(
		insertMessageSQL,
		UID,
		text,
		CreatedByUserID,
		createdAt,
		updatedAt,
	))
}

// updateUser sql query
var updateUserSQL = `UPDATE message
SET
	text = $2,
	updated_at = $3
WHERE message.id = $1
RETURNING message.id, message.uid, message.text, message.created_by_user_id, message.created_at, message.updated_at`

func execUpdateUser(
	db database.Transaction,
	ID entity.ID,
	text string,
	updatedAt time.Time,
) (sql.Result, error) {
	return db.Exec(
		updateUserSQL,
		ID,
		text,
		updatedAt,
	)
}