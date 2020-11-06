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
		if item, ok := result.(auth.User); ok {
			items[i] = &item
		} else {
			return nil, errCouldNotParse
		}
	}

	return items, nil
}

// FindByID get one message
func (f findRepository) FindByID(ID auth.ID) (*auth.User, error) {
	if result, ok := f.store.Get(ID.String()); ok {
		if s, ok := result.(userRecord); ok {
			return recordToUser(s), nil
		}
		return nil, errCouldNotParse
	}
	return nil, auth.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*auth.User, error) {
	for _, result := range f.store.List() {
		if s, ok := result.(userRecord); ok {
			if email == s.Email {
				return recordToUser(s), nil
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
		if s, ok := result.(userRecord); ok {
			if username == s.Username {
				return recordToUser(s), nil
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
