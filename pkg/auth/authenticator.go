package auth

import (
	"errors"
)

var (
	// ErrEmailRequired error decloration
	ErrEmailRequired = errors.New("email is required")
	// ErrPasswordRequired error decloration
	ErrPasswordRequired = errors.New("password is required")
)

// AuthenticateRepository authenticate user
type AuthenticateRepository interface {
	Authenticate(credentials Credentials, comparePasswords PasswordComparer) (bool, error)
}

// Authenticator authenticate with credentials.
type Authenticator struct {
	authenticateRepo AuthenticateRepository
}

type PasswordComparer = func(hashedPassword, password string) bool

func defaultComparePassword(hashedPassword, password string) bool {
	return ComparePasswords(hashedPassword, []byte(password))
}

// Authenticate by credentials
func (m *Authenticator) Authenticate(credentials Credentials) (bool, error) {

	if credentials.Email == "" {
		return false, ErrEmailRequired
	}

	if credentials.Password == "" {
		return false, ErrPasswordRequired
	}

	return m.authenticateRepo.Authenticate(Credentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	}, defaultComparePassword)
}

// NewAuthenticator returns a new Authenticator
func NewAuthenticator(a AuthenticateRepository) *Authenticator {
	return &Authenticator{a}
}
