package auth

import (
	"errors"

	"github.com/dwethmar/atami/pkg/model"
)

var (
	// ErrCouldNotDelete error
	ErrCouldNotDelete = errors.New("could not delete user")
)

// DeleterRepository deletes user
type DeleterRepository interface {
	Delete(ID model.UserID) error
}

// Deleter deletes messages.
type Deleter struct {
	deleteRepo DeleterRepository
}

// Delete a message
func (m *Deleter) Delete(ID model.UserID) error {
	return m.deleteRepo.Delete(ID)
}

// NewDeleter returns a new Listing
func NewDeleter(r DeleterRepository) *Deleter {
	return &Deleter{r}
}
