package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// searchRepository reads messages from memory
type searchRepository struct {
	store *memstore.Store
}

// ReadOne get one message
func (i searchRepository) Search(query user.Query) ([]*user.User, error) {
	s := make([]*user.User, 0)
	return s, nil
}

// NewSearcherRepository return a new in memory listin repository
func NewSearcherRepository(store *memstore.Store) user.SearchRepository {
	return &searchRepository{store}
}
