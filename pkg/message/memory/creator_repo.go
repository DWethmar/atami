package memory

import (
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// CreatorRepository stores new messages
type CreatorRepository struct {
	store *memstore.Store
	newID message.ID
}

// Create new message
func (i CreatorRepository) Create(newMessage message.NewMessage) (*message.Message, error) {
	i.newID++
	message := &message.Message{
		ID:        i.newID,
		Content:   newMessage.Content,
		CreatedAt: time.Now(),
	}
	i.store.Add(strconv.FormatInt(int64(message.ID), 10), message)
	return message, nil
}

// NewCreatorRepository creates new messages.
func NewCreatorRepository(store *memstore.Store) *CreatorRepository {
	return &CreatorRepository{
		store,
		0,
	}
}
