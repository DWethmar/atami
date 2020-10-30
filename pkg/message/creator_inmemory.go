package memory

import (
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

type inMemoryCreator struct {
	store *memstore.Store
	newID ID
}

func (i inMemoryCreator) Create(newMessage message.New) (Message, error) {
	i.newID++
	message := Message{
		ID:        i.newID,
		Content:   newMessage.Content,
		CreatedAt: time.Now(),
	}
	i.store.Add(strconv.FormatInt(int64(message.ID), 10), message)
	return message, nil
}

// NewInMemoryCreator creates new messages.
func NewInMemoryCreator(store *memstore.Store) *Creator {
	return NewCreator(
		inMemoryCreator{
			store,
		},
	)
}
