package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
)

// deleterRepository deletes messages from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one message
func (i deleterRepository) Delete(ID model.MessageID) error {
	if i.store.Delete(ID.String()) {
		return nil
	}
	return message.ErrCouldNotDelete
}

// NewDeleter return a new deleter
func NewDeleter(store *memstore.Store) *message.Deleter {
	return message.NewDeleter(&deleterRepository{store})
}
