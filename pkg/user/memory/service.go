package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

type service struct {
	user.Searcher
	user.Reader
	user.Deleter
	user.Creator
}

// NewService creates a new user service
func NewService(store *memstore.Store) user.Service {
	s := NewSearcher(store)
	r := NewReader(store)
	d := NewDeleter(store)
	c := NewCreator(store)

	return &service{
		Searcher: *s,
		Reader:   *r,
		Deleter:  *d,
		Creator:  *c,
	}
}
