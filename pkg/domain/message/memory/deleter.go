package memory

import (
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/memstore"
)

// deleterRepository deletes messages from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one message
func (i deleterRepository) Delete(ID int) error {
	messages := i.store.GetMessages()
	if messages.Delete(ID) {
		return nil
	}
	return message.ErrCouldNotDelete
}

// NewDeleter return a new deleter
func NewDeleter(store *memstore.Store) *message.Deleter {
	return message.NewDeleter(&deleterRepository{store})
}
