package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/auth"
)

type service struct {
	auth.Authenticator
	auth.Finder
	auth.Deleter
	auth.Registrator
	auth.Validator
}

// NewService creates a new user service
func NewService(db *sql.DB) auth.Service {
	var validator = auth.NewDefaultValidator()

	a := NewAuthenticator(db)
	f := NewFinder(db)
	d := NewDeleter(db)
	r := NewRegistrator(f, validator, db)

	return &service{
		Authenticator: *a,
		Finder:        *f,
		Deleter:       *d,
		Registrator:   *r,
		Validator:     *validator,
	}
}
