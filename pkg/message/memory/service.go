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

// New creates a new user service
func New(store *memstore.Store) *message.Service {
	r := NewFinder(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return message.NewService(
		*r,
		*d,
		*c,
	)
}
