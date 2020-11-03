package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// readerRepository reads messages from memory
type readerRepository struct {
	store *memstore.Store
}

// ReadOne get one message
func (i readerRepository) ReadOne(ID user.ID) (*user.User, error) {
	result, ok := i.store.Get(ID.String())
	if ok {
		user, ok := result.(user.User)
		if ok {
			return &user, nil
		}
		return nil, errors.New("error while parsing result")
	}
	return nil, user.ErrCouldNotFind
}

// ReadAll get multiple messages
func (i readerRepository) ReadAll() ([]*user.User, error) {
	results := i.store.List()
	items := make([]*user.User, len(results))

	for i, l := range results {
		if item, ok := l.(user.User); ok {
			items[i] = &item
		} else {
			return nil, errors.New("Error while parsing")
		}
	}

	return items, nil
}

// NewReaderRepository return a new in memory listin repository
func NewReaderRepository(store *memstore.Store) user.ReaderRepository {
	return &readerRepository{store}
}
