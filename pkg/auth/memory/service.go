package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

type service struct {
	auth.Authenticator
	auth.Finder
	auth.Deleter
	auth.Registrator
	auth.Validator
}

// New creates a new user service
func New(store *memstore.Store) *auth.Service {
	var validator = auth.NewDefaultValidator()

	a := NewAuthenticator(store)
	f := NewFinder(store)
	d := NewDeleter(store)
	r := NewRegistrator(f, validator, store)

	return auth.NewService(
		*a,
		*f,
		*d,
		*r,
		*validator,
	)
}
