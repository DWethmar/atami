package auth

import (
	"github.com/dwethmar/atami/pkg/user"
)

// Authenticator authenticate with credentials.
type Authenticator struct {
	finder *user.Finder
}

// Authenticate by credentials
func (m *Authenticator) Authenticate(credentials Credentials) (bool, error) {
	if credentials.Email == "" {
		return false, ErrEmailRequired
	}

	if credentials.Password == "" {
		return false, ErrPasswordRequired
	}

	usr, err := m.finder.FindByEmailWithPassword(credentials.Email)
	if err != nil {
		if err != user.ErrCouldNotFind {
			return false, err
		}
		return false, ErrAuthentication
	}

	if ComparePasswords(usr.Password, []byte(credentials.Password)) {
		return true, nil
	}

	return false, ErrAuthentication
}

// NewAuthenticator returns a new Authenticator
func NewAuthenticator(finder *user.Finder) *Authenticator {
	return &Authenticator{finder}
}
