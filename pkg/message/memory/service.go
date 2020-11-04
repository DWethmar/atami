package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

type service struct {
	message.Finder
	message.Deleter
	message.Creator
}

// NewService creates a new user service
func NewService(store *memstore.Store) message.Service {
	r := NewFinder(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return &service{
		Finder:  *r,
		Deleter: *d,
		Creator: *c,
	}
}
