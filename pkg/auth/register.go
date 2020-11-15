package auth

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

// RegisterRepository declares a storage repository
type RegisterRepository interface {
	Register(createUser HashedCreateUser) (*User, error)
}

// Registrator struct declaration
type Registrator struct {
	validator    *Validator
	finder       *Finder // TODO: refactor to repo
	registerRepo RegisterRepository
}

// Register registers a new user
func (m *Registrator) Register(newUser CreateUser) (*User, error) {
	if err := m.validator.ValidateNewUser(newUser); err != nil {
		return nil, err
	}

	if usr, err := m.finder.FindByEmail(newUser.Email); usr != nil && err == nil {
		return nil, ErrEmailAlreadyTaken
	} else if err != nil && err != ErrCouldNotFind {
		return nil, err
	}

	if usr, err := m.finder.FindByUsername(newUser.Username); usr != nil && err == nil {
		return nil, ErrUsernameAlreadyTaken
	} else if err != nil && err != ErrCouldNotFind {
		return nil, err
	}

	hashedPassword := HashPassword([]byte(newUser.Password))

	createUser := HashedCreateUser{
		Username:       newUser.Username,
		Email:          newUser.Email,
		HashedPassword: hashedPassword,
	}

	user, err := m.registerRepo.Register(createUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// NewRegistartor returns a new searcher
func NewRegistartor(r RegisterRepository, f *Finder, v *Validator) *Registrator {
	return &Registrator{
		registerRepo: r,
		finder:       f,
		validator:    v,
	}
}
