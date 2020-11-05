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
		if user, ok := result.(auth.User); ok {
			return &user, nil
		}
		return nil, errCouldNotParse
	}
	return nil, auth.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*auth.User, error) {
	for _, result := range f.store.List() {
		if item, ok := result.(auth.User); ok {
			if email == item.Email {
				return &item, nil
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
		if item, ok := result.(auth.User); ok {
			if username == item.Username {
				return &item, nil
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
