package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// readerRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// FindByID get one message
func (i findRepository) FindByID(ID message.ID) (*message.Message, error) {
	result, ok := i.store.Get(ID.String())
	if ok {
		if message, ok := result.(message.Message); ok {
			return &message, nil
		}
		return nil, errors.New("error while parsing result")
	}
	return nil, message.ErrCouldNotFind
}

// FindAll get multiple messages
func (i findRepository) FindAll() ([]*message.Message, error) {
	results := i.store.List()
	items := make([]*message.Message, len(results))

	for i, l := range results {
		if item, ok := l.(message.Message); ok {
			items[i] = &item
		} else {
			return nil, errors.New("Error while parsing")
		}
	}

	return items, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(store *memstore.Store) *message.Finder {
	return message.NewFinder(&findRepository{store})
}