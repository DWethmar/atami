package auth

import (
	"errors"
)

var (
	// ErrCouldNotDelete error
	ErrCouldNotDelete = errors.New("could not delete user")
)

// DeleterRepository deletes user
type DeleterRepository interface {
	Delete(ID int) error
}

// Deleter deletes messages.
type Deleter struct {
	deleteRepo DeleterRepository
}

// Delete a message
func (m *Deleter) Delete(ID int) error {
	return m.deleteRepo.Delete(ID)
}

// NewDeleter returns a new Listing
func NewDeleter(r DeleterRepository) *Deleter {
	return &Deleter{r}
}
