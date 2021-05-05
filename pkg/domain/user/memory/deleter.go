package memory

import (
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	store *memstore.Memstore
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	if i.store.GetUsers().Delete(ID) {
		return nil
	}
	return user.ErrCouldNotDelete
}

// NewDeleter return a new in deleter repo
func NewDeleter(store *memstore.Memstore) *user.Deleter {
	return user.NewDeleter(&deleterRepository{store})
}
