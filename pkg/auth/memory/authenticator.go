package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

// AuthenticatorRepository authenticates users by credentials.
type authenticatorRepository struct {
	store *memstore.Store
}

// Authenticate an user
func (a authenticatorRepository) Authenticate(credentials auth.Credentials, comparePasswords auth.PasswordComparer) (bool, error) {
	var rUser *userRecord

	for _, result := range a.store.List() {
		if record, ok := result.(userRecord); ok {
			if credentials.Email == record.Email {
				rUser = &record
			}
		} else {
			return false, errCouldNotParse
		}
	}

	if rUser == nil {
		return false, auth.ErrCouldNotFind
	}

	if comparePasswords(rUser.Password, credentials.Password) {
		return true, nil
	}

	return false, auth.ErrCouldNotFind
}

// NewAuthenticator return a new in authenticator
func NewAuthenticator(store *memstore.Store) *auth.Authenticator {
	return auth.NewAuthenticator(&authenticatorRepository{store})
}
