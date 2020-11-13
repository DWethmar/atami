package memory

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
	"github.com/segmentio/ksuid"
)

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID model.MessageID
}

// Create new message
func (i *creatorRepository) Create(newMessage message.NewMessage) (*message.Message, error) {
	i.newID++
	uid := model.MessageUID(ksuid.New().String())
	msg := message.Message{
		ID:        i.newID,
		UID:       uid,
		Text:      newMessage.Text,
		CreatedAt: time.Now(),
		CreatedBy: newMessage.CreatedBy,
	}
	i.store.Add(msg.ID.String(), msg)

	if value, ok := i.store.Get(msg.ID.String()); ok {
		if msg, ok := value.(message.Message); ok {
			return &msg, nil
		}
		return nil, errCouldNotParse
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
