package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/validate"
)

type service struct {
	user.Finder
	user.Deleter
	user.Creator
}

// NewService creates a new user service
func NewService(store *memstore.Store) user.Service {
	validator := user.NewValidator(validate.NewEmailValidator())

	f := NewFinder(store)
	d := NewDeleter(store)
	c := NewCreator(validator, store)

	return &service{
		Finder:  *f,
		Deleter: *d,
		Creator: *c,
	}
}
