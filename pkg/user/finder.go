package user

import "errors"

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find message")
)

// FindRepository defines a messsage listing repository
type FindRepository interface {
	FindAll() ([]*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(ID ID) (*User, error)
}

// Finder searches messages.
type Finder struct {
	findRepo FindRepository
}

// FindAll return a list of list items.
func (m *Finder) FindAll() ([]*User, error) {
	return m.findRepo.FindAll()
}

// FindByID return a list of list items.
func (m *Finder) FindByID(ID ID) (*User, error) {
	return m.findRepo.FindByID(ID)
}

// FindByEmail searches users by email
func (m *Finder) FindByEmail(email string) (*User, error) {
	return m.findRepo.FindByEmail(email)
}

// NewFinder returns a new searcher
func NewFinder(r FindRepository) *Finder {
	return &Finder{r}
}
