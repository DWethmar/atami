package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
	"github.com/segmentio/ksuid"
)

var insertUser = fmt.Sprintf(`
INSERT INTO %s (
	uid,
	text, 
	created_by_user_id,
	created_on
)
VALUES ($1, $2, $3, $4) RETURNING id`, tableName)

// creatorRepository stores new messages
type creatorRepository struct {
	db *sql.DB
}

// Create new message
func (i creatorRepository) Create(newMessage message.NewMessage) (*message.Message, error) {

	stmt, err := i.db.Prepare(insertUser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var messageID int
	if stmt.QueryRow(
		model.MessageUID(ksuid.New().String()),
		newMessage.Text,
		newMessage.CreatedBy,
		time.Now().UTC(),
	).Scan(&messageID); err != nil {
		return nil, err
	}

	if messageID != 0 {
		entry := &message.Message{}
		if err := i.db.QueryRow(getMessageByID, messageID).Scan(
			&entry.ID,
			&entry.UID,
			&entry.Text,
			&entry.CreatedBy,
			&entry.CreatedAt,
		); err != nil {
			return nil, err
		}
		return entry, nil
	}

	return nil, fmt.Errorf("could not create message with id %v", messageID)
}

// NewCreator creates new messages creator.
func NewCreator(db *sql.DB) *message.Creator {
	return message.NewCreator(&creatorRepository{
		db,
	})
}
