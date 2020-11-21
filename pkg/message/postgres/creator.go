package postgres

import (
	"database/sql"
	"time"

	"github.com/dwethmar/atami/pkg/message"
	"github.com/segmentio/ksuid"
)

// creatorRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

// Create new message
func (i creatorRepository) Create(newMessage message.CreateMessage) (*message.Message, error) {
	msg, err := queryRowInsertMessage(
		i.db,
		ksuid.New().String(),
		newMessage.Text,
		newMessage.CreatedByUserID,
		time.Now().UTC(),
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
