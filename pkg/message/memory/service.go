package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

type service struct {
	message.Reader
	message.Deleter
	message.Creator
}

// NewService creates a new user service
func NewService(store *memstore.Store) message.Service {
	r := NewReader(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return &service{
		Reader:  *r,
		Deleter: *d,
		Creator: *c,
	}
}
