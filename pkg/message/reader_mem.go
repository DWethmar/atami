package message

import (
	"errors"
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
)

// MemReader reads messages from memory
type MemReader struct {
	store *memstore.Store
}

// ReadOne get one message
func (i MemReader) ReadOne(ID ID) (*Message, error) {
	result, ok := i.store.Get(strconv.FormatInt(int64(ID), 10))
	if ok {
		message, ok := result.(Message)
		if ok {
			return &message, nil
		}
		return nil, errors.New("error while parsing result")
	}
	return nil, ErrCouldNotFind
}

// ReadAll get multiple messages
func (i MemReader) ReadAll() ([]*Message, error) {
	results := i.store.List()
	items := make([]*Message, len(results))

	for i, l := range results {
		if item, ok := l.(Message); ok {
			items[i] = &item
		} else {
			return nil, errors.New("Error while parsing")
		}
	}

	return items, nil
}

// NewMemReader return a new in memory listin repository
func NewMemReader(store *memstore.Store) *Reader {
	return &Reader{MemReader{store}}
}
