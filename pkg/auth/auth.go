package auth

import (
	"errors"

	"github.com/dwethmar/atami/pkg/user"
)

var (
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("Email already taken")
)

// NewUser struct decloration
type NewUser struct {
	Username string
	Email    string
	Password string
}

// CreateUser creates a new User
func CreateUser(creator user.Creator, reader user.Reader, newUser NewUser) *user.User {
	return &user.User{}
}
