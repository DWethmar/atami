package user

import "errors"

var (
	// ErrUnknownOperator error
	ErrUnknownOperator = errors.New("Unknown operator error")
	// ErrUnknownField error
	ErrUnknownField = errors.New("Unknown field error")
	// ErrInvalidValue error
	ErrInvalidValue = errors.New("Invalid value error")
)

// SearchRepository defines a messsage listing repository
type SearchRepository interface {
	SearchByEmail(email string) ([]*User, error)
}

// Searcher searches messages.
type Searcher struct {
	searchRepo SearchRepository
}

// SearchByEmail searches users by email
func (m *Searcher) SearchByEmail(email string) ([]*User, error) {
	return m.searchRepo.SearchByEmail(email)
}

// NewSearcher returns a new searcher
func NewSearcher(r SearchRepository) *Searcher {
	return &Searcher{r}
}
