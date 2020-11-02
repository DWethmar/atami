package user

// SearchRepository defines a messsage listing repository
type SearchRepository interface {
	Search(query Query) ([]*User, error)
}

// Searcher searches messages.
type Searcher struct {
	searchRepo SearchRepository
}

// Search return a list of list items based on a query.
func (m *Searcher) Search(query Query) ([]*User, error) {
	return m.searchRepo.Search(query)
}

// NewSearcher returns a new searcher
func NewSearcher(r SearchRepository) *Searcher {
	return &Searcher{r}
}
