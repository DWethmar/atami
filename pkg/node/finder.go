package node

import (
	"errors"
)

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find node")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	FindByUID(UID string) (*Node, error)
	FindByID(ID int) (*Node, error)
	Find(limit, offset int) ([]*Node, error)
}

// Finder lists nodes.
type Finder struct {
	readerRepo FindRepository
}

// FindByUID return a list of list items.
func (m *Finder) FindByUID(UID string) (*Node, error) {
	return m.readerRepo.FindByUID(UID)
}

// FindByID return a list of list items.
func (m *Finder) FindByID(ID int) (*Node, error) {
	return m.readerRepo.FindByID(ID)
}

// Find return a list of list items.
func (m *Finder) Find(page, size int) ([]*Node, error) {
	return m.readerRepo.Find(size, page*size)
}

// NewFinder returns a new Listing
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
