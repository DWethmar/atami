package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/validate"
)

type service struct {
	auth.Finder
	auth.Deleter
	auth.Registrator
}

// NewService creates a new user service
func NewService(store *memstore.Store) auth.Service {
	var validator = auth.NewValidator(
		validate.NewUsernameValidator(),
		validate.NewEmailValidator(),
		validate.NewPasswordValidator(),
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
