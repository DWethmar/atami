package auth

import (
	"errors"
)

var (
	// ErrUsernameAlreadyTaken error decloration
	ErrUsernameAlreadyTaken = errors.New("username is unavailable")
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("username is unavailable")
	// ErrPwdNotSet error decloration
	ErrPwdNotSet = errors.New("password not set")
)

// RegisterRepository declares a storage repository
type RegisterRepository interface {
	Register(createUser CreateUser) (*User, error)
}

// Registrator struct declaration
type Registrator struct {
	validator    *Validator
	finder       *Finder
	registerRepo RegisterRepository
}

// Register registers a new user
func (m *Registrator) Register(newUser RegisterUser) (*User, error) {
	if usr, err := m.finder.FindByEmail(newUser.Email); usr != nil && err == nil {
		return nil, ErrEmailAlreadyTaken
	} else if err != ErrCouldNotFind {
		return nil, err
	}

	if usr, err := m.finder.FindByUsername(newUser.Username); usr != nil && err == nil {
		return nil, ErrUsernameAlreadyTaken
	} else if err != ErrCouldNotFind {
		return nil, err
	}

	if err := m.validator.ValidateNewUser(newUser); err != nil {
		return nil, err
	}

	createUser := CreateUser{
		Username:       newUser.Username,
		Email:          newUser.Email,
		HashedPassword: HashPassword([]byte(newUser.PlainPassword)),
	}

	return m.registerRepo.Register(createUser)
}

// NewRegistartor returns a new searcher
func NewRegistartor(r RegisterRepository, f *Finder, v *Validator) *Registrator {
	return &Registrator{
		registerRepo: r,
		finder:       f,
		validator:    v,
	}
}
