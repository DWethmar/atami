package memory

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/segmentio/ksuid"
)

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID message.ID
}

// Create new message
func (i creatorRepository) Create(newMessage message.NewMessage) (*message.Message, error) {
	i.newID++
	uid := message.UID(ksuid.New().String())
	i.store.Add(string(uid), message.Message{
		ID:        i.newID,
		UID:       uid,
		Text:      newMessage.Content,
		CreatedAt: time.Now(),
	})

	if value, ok := i.store.Get(string(uid)); ok {
		if msg, ok := value.(message.Message); ok {
			return &msg, nil
		}
		return nil, errors.New("Error parsing message")
	}

	return nil, errors.New("Could not find message")
}

// NewCreator creates new messages creator.
func NewCreator(store *memstore.Store) *message.Creator {
	return message.NewCreator(&creatorRepository{
		store,
		0,
	})
}
