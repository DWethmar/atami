package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dwethmar/atami/pkg/auth"
)

type authUser struct {
	auth.User
	Password string
}

var authUserByEmail = `
SELECT
	password,
	email
FROM public.user
WHERE email = $1
LIMIT 1`

// AuthenticatorRepository authenticates users by credentials.
type authenticatorRepository struct {
	db *sql.DB
}

// Authenticate an user
func (a authenticatorRepository) Authenticate(credentials auth.Credentials, comparePasswords auth.PasswordComparer) (bool, error) {
	entry := &authUser{}

	fmt.Printf("Yo whats is going on with those emails %v \n", credentials.Email)

	if err := a.db.QueryRow(authUserByEmail, credentials.Email).Scan(
		&entry.Password,
		&entry.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return false, auth.ErrCouldNotFind
		}

		fmt.Printf(" :( %v", err)

		return false, err
	}

	return comparePasswords(entry.Password, credentials.Password), nil
}

// NewAuthenticator return a new in authenticator
func NewAuthenticator(db *sql.DB) *auth.Authenticator {
	return auth.NewAuthenticator(&authenticatorRepository{db})
}
