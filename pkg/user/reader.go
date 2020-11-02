package user

import "errors"

var (
	// ErrCouldNotFind error
	ErrCouldNotFind = errors.New("could not find user")
)

// ReaderRepository defines a messsage listing repository
type ReaderRepository interface {
	ReadOne(ID ID) (*User, error)
	ReadAll() ([]*User, error)
}

// Reader lists messages.
type Reader struct {
	readerRepo ReaderRepository
}

// ReadOne return a list of list items.
func (m *Reader) ReadOne(ID ID) (*User, error) {
	return m.readerRepo.ReadOne(ID)
}

// ReadAll return a list of list items.
func (m *Reader) ReadAll() ([]*User, error) {
	return m.readerRepo.ReadAll()
}

// NewReader returns a new Listing
func NewReader(r ReaderRepository) *Reader {
	return &Reader{r}
}
