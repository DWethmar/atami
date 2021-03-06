package user

import (
	"errors"
)

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find user")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	Find() ([]*User, error)
	FindByEmail(email string) (*User, error)
	FindByEmailWithPassword(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByID(ID int) (*User, error)
	FindByUID(UID string) (*User, error)
}

// Finder searches messages. // TODO: rename to reader
type Finder struct {
	findRepo FindRepository
}

// Find return a list of list items.
func (m *Finder) Find() ([]*User, error) {
	return m.findRepo.Find()
}

// FindByEmail search for user with email
func (m *Finder) FindByEmail(email string) (*User, error) {
	return m.findRepo.FindByEmail(email)
}

// FindByEmailWithPassword search for user with email and set password
func (m *Finder) FindByEmailWithPassword(email string) (*User, error) {
	return m.findRepo.FindByEmailWithPassword(email)
}

// FindByUsername search for user with username
func (m *Finder) FindByUsername(username string) (*User, error) {
	return m.findRepo.FindByUsername(username)
}

// FindByID search for user with provided ID
func (m *Finder) FindByID(ID int) (*User, error) {
	return m.findRepo.FindByID(ID)
}

// FindByUID search for user with provided UID
func (m *Finder) FindByUID(UID string) (*User, error) {
	return m.findRepo.FindByUID(UID)
}

// NewFinder returns a new searcher
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
