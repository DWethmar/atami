package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/validate"
)

type service struct {
	user.Finder
	user.Deleter
	user.Registrator
}

// NewService creates a new user service
func NewService(store *memstore.Store) user.Service {
	var validator = user.NewValidator(
		validate.NewUsernameValidator(),
		validate.NewEmailValidator(),
	)

	f := NewFinder(store)
	d := NewDeleter(store)
	r := NewRegistrator(f, validator, store)

	return &service{
		Finder:      *f,
		Deleter:     *d,
		Registrator: *r,
	}
}
