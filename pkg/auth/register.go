package auth

import (
	"errors"

	"github.com/dwethmar/atami/pkg/user"
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
	Register(createUser CreateUser) (*user.User, error)
}

// Registrator struct declaration
type Registrator struct {
	validator *Validator
	finder    *user.Finder
	creator   *user.Creator
}

// Register registers a new user
func (m *Registrator) Register(newUser CreateUser) (*user.User, error) {
	if err := m.validator.ValidateNewUser(newUser); err != nil {
		return nil, err
	}

	if usr, err := m.finder.FindByEmail(newUser.Email); usr != nil && err == nil {
		return nil, ErrEmailAlreadyTaken
	} else if err != nil && err != user.ErrCouldNotFind {
		return nil, err
	}

	if usr, err := m.finder.FindByUsername(newUser.Username); usr != nil && err == nil {
		return nil, ErrUsernameAlreadyTaken
	} else if err != nil && err != user.ErrCouldNotFind {
		return nil, err
	}

	hashedPassword := HashPassword([]byte(newUser.Password))

	createUser := user.CreateUser{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashedPassword,
	}

	user, err := m.creator.Create(createUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// NewRegistrator returns a new searcher
func NewRegistrator(r *user.Creator, f *user.Finder) *Registrator {
	return &Registrator{
		creator:   r,
		finder:    f,
		validator: NewDefaultValidator(),
	}
}
