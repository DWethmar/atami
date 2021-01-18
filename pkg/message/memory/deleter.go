package memory

import (
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// deleterRepository deletes messages from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one message
func (i deleterRepository) Delete(ID int) error {
	messages := i.store.GetMessages()
	if messages.Delete(strconv.Itoa(ID)) {
		return nil
	}
	return message.ErrCouldNotDelete
}

// NewDeleter return a new deleter
func NewDeleter(store *memstore.Store) *message.Deleter {
	return message.NewDeleter(&deleterRepository{store})
}
