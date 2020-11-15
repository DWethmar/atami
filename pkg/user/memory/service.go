package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// New creates a new user service
func New(store *memstore.Store) *user.Service {
	var validator = user.NewDefaultValidator()

	f := NewFinder(store)
	d := NewDeleter(store)
	r := NewCreator(validator, store)

	return user.NewService(
		*f,
		*d,
		*r,
		*validator,
	)
}
