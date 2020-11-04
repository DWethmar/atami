package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

type service struct {
	user.Finder
	user.Deleter
	user.Creator
}

// NewService creates a new user service
func NewService(store *memstore.Store) user.Service {
	f := NewFinder(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return &service{
		Finder:  *f,
		Deleter: *d,
		Creator: *c,
	}
}
