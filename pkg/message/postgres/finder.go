package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dwethmar/atami/pkg/message"
)

var getMessageByID = fmt.Sprintf(`
SELECT
	id,
	uid,
	text, 
	created_by_user_id,
	created_at
FROM  %s 
WHERE id = $1`, Table)

// findRepository reads messages from memory
type findRepository struct {
	db *sql.DB
}

// FindAll get multiple messages
func (i *findRepository) Find() ([]*message.Message, error) {
	rows, err := i.db.Query(getMessages, 100, 1)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]*message.Message, 0)
	for rows.Next() {
		entry := &message.Message{}

		if err := rows.Scan(
			&entry.ID,
			&entry.UID,
			&entry.Text,
			&entry.CreatedByUserID,
			&entry.CreatedAt,
		); err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

// FindByID get one message
func (i *findRepository) FindByID(ID int) (*message.Message, error) {
	entry := &message.Message{}
	if err := i.db.QueryRow(getMessageByID, ID).Scan(
		&entry.ID,
		&entry.UID,
		&entry.Text,
		&entry.CreatedByUserID,
		&entry.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, message.ErrCouldNotFind
		}
		return nil, err
	}
	return entry, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(db *sql.DB) *message.Finder {
	return message.NewFinder(&findRepository{db})
}
