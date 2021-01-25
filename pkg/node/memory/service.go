package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
)

type service struct {
	node.Finder
	node.Deleter
	node.Creator
}

// New creates a new user service
func New(store *memstore.Store) *node.Service {
	r := NewFinder(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return node.NewService(
		*r,
		*d,
		*c,
	)
}
