package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// searchRepository reads messages from memory
type searchRepository struct {
	store *memstore.Store
}

func (s *searchRepository) SearchByEmail(email string) ([]*user.User, error) {
	results := s.store.List()
	items := make([]*user.User, 0)

	for _, l := range results {
		if item, ok := l.(user.User); ok {
			if email == item.Email {
				items = append(items, &item)
			}
		} else {
			return nil, errors.New("Error while parsing")
		}
	}

	return items, nil
}

// NewSearcher return a new in memory listin repository
func NewSearcher(store *memstore.Store) *user.Searcher {
	return user.NewSearcher(&searchRepository{store})
}
