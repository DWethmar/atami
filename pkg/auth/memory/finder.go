package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// FindAll get multiple messages
func (f findRepository) FindAll() ([]*auth.User, error) {
	results := f.store.List()
	items := make([]*auth.User, len(results))

	for i, result := range results {
		if record, ok := result.(userRecord); ok {
			items[i] = recordToUser(record)
		} else {
			return nil, errCouldNotParse
		}
	}

	return items, nil
}

// FindByID get one message
func (f findRepository) FindByID(ID auth.ID) (*auth.User, error) {
	if result, ok := f.store.Get(ID.String()); ok {
		if record, ok := result.(userRecord); ok {
			return recordToUser(record), nil
		}
		return nil, errCouldNotParse
	}
	return nil, auth.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*auth.User, error) {
	for _, result := range f.store.List() {
		if record, ok := result.(userRecord); ok {
			if email == record.Email {
				return recordToUser(record), nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}

	return nil, auth.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*auth.User, error) {
	for _, result := range f.store.List() {
		if record, ok := result.(userRecord); ok {
			if username == record.Username {
				return recordToUser(record), nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}

	return nil, auth.ErrCouldNotFind
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *auth.Finder {
	return auth.NewFinder(&findRepository{store})
}
