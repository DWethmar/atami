package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID int
}

// Create new message
func (i *creatorRepository) Create(newMsg message.CreateMessage) (*message.Message, error) {
	messages := i.store.GetMessages()
	users := i.store.GetUsers()

	i.newID++
	msg := message.Message{
		ID:              i.newID,
		UID:             newMsg.UID,
		Text:            newMsg.Text,
		CreatedByUserID: newMsg.CreatedByUserID,
		CreatedAt:       newMsg.CreatedAt,
	}

	if _, ok := users.Get(msg.CreatedByUserID); !ok {
		return nil, errors.New("user not found")
	}

	messages.Put(msg.ID, util.ToMemory(msg))
	if r, ok := messages.Get(msg.ID); ok {
		msg := util.FromMemory(r)
		return &msg, nil
	}

	return nil, errors.New("could not find message")
}

// NewCreator creates new messages creator.
func NewCreator(store *memstore.Store) *message.Creator {
	return message.NewCreator(&creatorRepository{
		store,
		0,
	})
}
