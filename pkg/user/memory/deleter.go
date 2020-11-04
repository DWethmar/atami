package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one user
func (i deleterRepository) Delete(ID user.ID) error {
	if i.store.Delete(ID.String()) {
		return nil
	}
	return user.ErrCouldNotDelete
}

// NewDeleter return a new in deleter repo
func NewDeleter(store *memstore.Store) *user.Deleter {
	return user.NewDeleter(&deleterRepository{store})
}