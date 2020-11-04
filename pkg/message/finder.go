package message

import "errors"

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find message")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	FindByID(ID ID) (*Message, error)
	FindAll() ([]*Message, error)
}

// Finder lists messages.
type Finder struct {
	readerRepo FindRepository
}

// FindByID return a list of list items.
func (m *Finder) FindByID(ID ID) (*Message, error) {
	return m.readerRepo.FindByID(ID)
}

// FindAll return a list of list items.
func (m *Finder) FindAll() ([]*Message, error) {
	return m.readerRepo.FindAll()
}

// NewFinder returns a new Listing
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
