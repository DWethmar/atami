package memory

import (
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/memstore"
)

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
