package user

import (
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

var (
	// ErrUsernameAlreadyTaken error decloration
	ErrUsernameAlreadyTaken = errors.New("username is unavailable")
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("email is unavailable")
)

// CreateRepository declares a storage repository
type CreateRepository interface {
	Create(createUser CreateUser) (*User, error)
}

// Creator struct declaration
type Creator struct {
	validator  *Validator
	createRepo CreateRepository
	finder     *Finder
}

// Create registers a new user
func (m *Creator) Create(createUser CreateUser) (*User, error) {
	if err := m.validator.ValidateCreateUser(createUser); err != nil {
		return nil, err
	}

	if usr, err := m.finder.FindByEmail(createUser.Email); usr != nil && err == nil {
		return nil, ErrEmailAlreadyTaken
	} else if err != nil && err != ErrCouldNotFind {
		return nil, err
	}

	if usr, err := m.finder.FindByUsername(createUser.Username); usr != nil && err == nil {
		return nil, ErrUsernameAlreadyTaken
	} else if err != nil && err != ErrCouldNotFind {
		return nil, err
	}

	hashedPassword := HashPassword([]byte(createUser.Password))

	return m.createRepo.Create(CreateUser{
		UID:       ksuid.New().String(),
		Username:  createUser.Username,
		Email:     createUser.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
}

// NewCreator returns a new searcher
func NewCreator(r CreateRepository, finder *Finder) *Creator {
	return &Creator{
		createRepo: r,
		validator:  NewValidator(),
		finder:     finder,
	}
}
