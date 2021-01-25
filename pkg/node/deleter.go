package node

import (
	"errors"
)

var (
	// ErrCouldNotDelete error
	ErrCouldNotDelete = errors.New("could not delete node")
)

// DeleterRepository deletes messsages
type DeleterRepository interface {
	Delete(ID int) error
}

// Deleter deletes nodes.
type Deleter struct {
	deleteRepo DeleterRepository
}

// Delete a node
func (m *Deleter) Delete(ID int) error {
	return m.deleteRepo.Delete(ID)
}

// NewDeleter returns a new Listing
func NewDeleter(r DeleterRepository) *Deleter {
	return &Deleter{r}
}
