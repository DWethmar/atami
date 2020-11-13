package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/model"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	store *memstore.Store
}

// Delete deletes one user
func (i deleterRepository) Delete(ID model.UserID) error {
	if i.store.Delete(ID.String()) {
		return nil
	}
	return auth.ErrCouldNotDelete
}

// NewDeleter return a new in deleter repo
func NewDeleter(store *memstore.Store) *auth.Deleter {
	return auth.NewDeleter(&deleterRepository{store})
}
