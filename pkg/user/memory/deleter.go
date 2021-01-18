package memory

import (
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	if i.store.GetUsers().Delete(strconv.Itoa(ID)) {
		return nil
	}
	return user.ErrCouldNotDelete
}

// NewDeleter return a new in deleter repo
func NewDeleter(store *memstore.Store) *user.Deleter {
	return user.NewDeleter(&deleterRepository{store})
}
