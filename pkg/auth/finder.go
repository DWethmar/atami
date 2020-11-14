package auth

import (
	"errors"

	"github.com/dwethmar/atami/pkg/model"
)

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find user")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	Find() ([]*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByID(ID model.UserID) (*User, error)
	FindByUID(UID model.UserUID) (*User, error)
}

// Finder searches messages.
type Finder struct {
	findRepo FindRepository
}

// Find return a list of list items.
func (m *Finder) Find() ([]*model.User, error) {
	results, err := m.findRepo.Find()
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, len(results))
	for i, result := range results {
		users[i] = toUser(result)
	}

	return users, nil
}

// FindByEmail search for user with email
func (m *Finder) FindByEmail(email string) (*model.User, error) {
	user, err := m.findRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return toUser(user), nil
}

// FindByUsername search for user with username
func (m *Finder) FindByUsername(username string) (*model.User, error) {
	user, err := m.findRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return toUser(user), nil
}

// FindByID search for user with provided ID
func (m *Finder) FindByID(ID model.UserID) (*model.User, error) {
	user, err := m.findRepo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return toUser(user), nil
}

// FindByUID search for user with provided UID
func (m *Finder) FindByUID(UID model.UserUID) (*model.User, error) {
	user, err := m.findRepo.FindByUID(UID)
	if err != nil {
		return nil, err
	}
	return toUser(user), nil
}

// NewFinder returns a new searcher
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
