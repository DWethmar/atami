package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
)

// deleterRepository deletes nodes from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one node
func (i deleterRepository) Delete(ID int) error {
	nodes := i.store.GetNodes()
	if nodes.Delete(ID) {
		return nil
	}
	return node.ErrCouldNotDelete
}

// NewDeleter return a new deleter
func NewDeleter(store *memstore.Store) *node.Deleter {
	return node.NewDeleter(&deleterRepository{store})
}
