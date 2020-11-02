package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// deleterRepository deletes messages from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one message
func (i deleterRepository) Delete(ID message.ID) error {
	if i.store.Delete(ID.String()) {
		return nil
	}
	return errors.New("Something went wrong while deleting a message from memory")
}

// NewDeleterRepository return a new in deleter repo
func NewDeleterRepository(store *memstore.Store) message.DeleterRepository {
	return &deleterRepository{store}
}
