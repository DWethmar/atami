package memory

import (
	"errors"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// readerRepository reads messages from memory
type readerRepository struct {
	store *memstore.Store
}

// ReadOne get one message
func (i readerRepository) ReadOne(ID message.ID) (*message.Message, error) {
	result, ok := i.store.Get(ID.String())
	if ok {
		message, ok := result.(message.Message)
		if ok {
			return &message, nil
		}
		return nil, errors.New("error while parsing result")
	}
	return nil, message.ErrCouldNotFind
}

// ReadAll get multiple messages
func (i readerRepository) ReadAll() ([]*message.Message, error) {
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

// NewReader return a new in memory listin reader
func NewReader(store *memstore.Store) *message.Reader {
	return message.NewReader(&readerRepository{store})
}
