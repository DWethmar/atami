package message

import (
	"errors"
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
)

type inMemoryListing struct {
	store *memstore.Store
}

func (i inMemoryListing) ReadOne(ID ID) (*Message, error) {
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

func (i inMemoryListing) ReadAll() ([]*Message, error) {
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

// NewInMemoryReader return a new inmemory listin repository
func NewInMemoryReader(store *memstore.Store) *Reader {
	return &Reader{inMemoryListing{store}}
}
