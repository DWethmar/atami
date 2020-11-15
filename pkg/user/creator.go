package user

import (
	"errors"
)

var (
	// ErrUsernameAlreadyTaken error decloration
	ErrUsernameAlreadyTaken = errors.New("username is unavailable")
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("email is unavailable")
	// ErrPwdNotSet error decloration
	ErrPwdNotSet = errors.New("password not set")
)

// CreateRepository declares a storage repository
type CreateRepository interface {
	Create(createUser CreateUser) (*User, error)
}

// Creator struct declaration
type Creator struct {
	validator  *Validator
	createRepo CreateRepository
}

// Create registers a new user
func (m *Creator) Create(createUser CreateUser) (*User, error) {
	if err := m.validator.ValidateCreateUser(createUser); err != nil {
		return nil, err
	}
	return m.createRepo.Create(createUser)
}

// NewCreator returns a new searcher
func NewCreator(r CreateRepository, v *Validator) *Creator {
	return &Creator{
		createRepo: r,
		validator:  v,
	}
}
