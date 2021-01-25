package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/message"
)

// creatorRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

// Create new message
func (i creatorRepository) Create(newMsg message.CreateAction) (*message.Message, error) {
	msg, err := queryRowInsertMessage(
		i.db,
		newMsg.UID,
		newMsg.Text,
		newMsg.CreatedByUserID,
		newMsg.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// NewCreator creates new messages creator.
func NewCreator(db *sql.DB) *message.Creator {
	return message.NewCreator(&creatorRepository{
		db,
	})
}
