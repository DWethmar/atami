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

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage NewUser) (*User, error)
}

// Creator creates messages.
type Creator struct {
	validator  *Validator
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newUser NewUser) (*User, error) {
	if err := m.validator.ValidateNewUser(newUser); err != nil {
		return nil, err
	}
	return m.createRepo.Create(newUser)
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository, v *Validator) *Creator {
	return &Creator{v, r}
}
