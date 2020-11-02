package user

import (
	"errors"
)

var (
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("email is already taken")
	// ErrPwdNotSet error decloration
	ErrPwdNotSet = errors.New("password not set")
)

// NewUser struct declaration
type NewUser struct {
	Username string
	Email    string
	Password string
}

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage NewUser) (*User, error)
}

// Creator creates messages.
type Creator struct {
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newUser NewUser) (*User, error) {
	return m.createRepo.Create(newUser)
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{r}
}
