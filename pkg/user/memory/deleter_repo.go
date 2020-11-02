package memory

import (
	"errors"

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
	return errors.New("Something went wrong while deleting a user from memory")
}

// NewDeleterRepository return a new in deleter repo
func NewDeleterRepository(store *memstore.Store) user.DeleterRepository {
	return &deleterRepository{store}
}
