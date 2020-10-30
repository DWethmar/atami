package message

import (
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
)

type MemoryCreator struct {
	store *memstore.Store
	newID ID
}

func (i MemoryCreator) Create(newMessage NewMessage) (*Message, error) {
	i.newID++
	message := Message{
		ID:        i.newID,
		Content:   newMessage.Content,
		CreatedAt: time.Now(),
	}
	i.store.Add(strconv.FormatInt(int64(message.ID), 10), message)
	return &message, nil
}

// NewInMemoryCreator creates new messages.
func NewMemCreator(store *memstore.Store) *Creator {
	return NewCreator(
		MemoryCreator{
			store,
			0,
		},
	)
}
