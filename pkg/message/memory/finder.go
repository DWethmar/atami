package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
)

// readerRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// FindByID get one message
func (i *findRepository) FindByID(ID model.MessageID) (*message.Message, error) {
	result, ok := i.store.Get(ID.String())
	if ok {
		if message, ok := result.(message.Message); ok {
			return &message, nil
		}
		return nil, errCouldNotParse
	}
	return nil, message.ErrCouldNotFind
}

// FindAll get multiple messages
func (i *findRepository) FindAll() ([]*message.Message, error) {
	results := i.store.List()
	items := make([]*message.Message, len(results))

	for i, l := range results {
		if item, ok := l.(message.Message); ok {
			items[i] = &item
		} else {
			return nil, errCouldNotParse
		}
	}

	return items, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(store *memstore.Store) *message.Finder {
	return message.NewFinder(&findRepository{store})
}
