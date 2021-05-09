package postgres

import (
	"time"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
)

// creatorRepository stores new messages
type creatorRepository struct {
	db database.Transaction
}

// Create new message
func (i creatorRepository) Create(newMsg message.CreateMessage) (*message.Message, error) {
	msg, err := queryRowInsertMessage(
		i.db,
		newMsg.UID,
		newMsg.Text,
		newMsg.CreatedByUserID,
		newMsg.CreatedAt,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// NewCreator creates new messages creator.
func NewCreator(db database.Transaction) *message.Creator {
	return message.NewCreator(&creatorRepository{
		db,
	})
}
