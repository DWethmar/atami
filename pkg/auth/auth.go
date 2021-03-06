package auth

import "errors"

var (
	// ErrAuthentication error decloration
	ErrAuthentication = errors.New("failure to authenticate")
	// ErrEmailRequired error decloration
	ErrEmailRequired = errors.New("email is required")
	// ErrPasswordRequired error decloration
	ErrPasswordRequired = errors.New("password is required")
)

// Credentials is information used to authenticate an user
type Credentials struct {
	Email    string
	Password string
}
