package memory

import (
	"errors"
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

// ReaderRepository reads messages from memory
type ReaderRepository struct {
	store *memstore.Store
}

// ReadOne get one message
func (i ReaderRepository) ReadOne(ID message.ID) (*message.Message, error) {
	result, ok := i.store.Get(strconv.FormatInt(int64(ID), 10))
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
func (i ReaderRepository) ReadAll() ([]*message.Message, error) {
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

// NewReaderRepository return a new in memory listin repository
func NewReaderRepository(store *memstore.Store) *ReaderRepository {
	return &ReaderRepository{store}
}
