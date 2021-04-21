package message

import (
	"time"

	"github.com/segmentio/ksuid"
)

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage CreateMessage) (*Message, error) // return int
}

// Creator creates messages.
type Creator struct {
	validator  *Validator
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(createMessage CreateMessage) (*Message, error) {
	// TODO validate!
	return m.createRepo.Create(CreateMessage{
		UID:             ksuid.New().String(),
		Text:            createMessage.Text,
		CreatedByUserID: createMessage.CreatedByUserID,
		CreatedAt:       time.Now().UTC(),
	})
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{
		NewDefaultValidator(),
		r,
	}
}
