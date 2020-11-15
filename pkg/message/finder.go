package message

import (
	"errors"
)

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find message")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	FindByID(ID int) (*Message, error)
	Find() ([]*Message, error)
}

// Finder lists messages.
type Finder struct {
	readerRepo FindRepository
}

// FindByID return a list of list items.
func (m *Finder) FindByID(ID int) (*Message, error) {
	return m.readerRepo.FindByID(ID)
}

// Find return a list of list items.
func (m *Finder) Find() ([]*Message, error) {
	return m.readerRepo.Find()
}

// NewFinder returns a new Listing
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
