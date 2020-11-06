package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

// authenticateRepository authenticates user from memory
type authenticateRepository struct {
	store *memstore.Store
}

// Delete deletes one user
func (a authenticateRepository) Authenticate(credentials auth.Credentials) (bool, error) {
	return false, auth.ErrCouldNotDelete
}

// NewAuthenticator return a new in authenticator
func NewAuthenticator(store *memstore.Store) *auth.Deleter {
	return auth.NewDeleter(&deleterRepository{store})
}
