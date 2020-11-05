package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// FindAll get multiple messages
func (f findRepository) FindAll() ([]*user.User, error) {
	results := f.store.List()
	items := make([]*user.User, len(results))

	for i, result := range results {
		if item, ok := result.(user.User); ok {
			items[i] = &item
		} else {
			return nil, errCouldNotParse
		}
	}

	return items, nil
}

// FindByID get one message
func (f findRepository) FindByID(ID user.ID) (*user.User, error) {
	if result, ok := f.store.Get(ID.String()); ok {
		if user, ok := result.(user.User); ok {
			return &user, nil
		}
		return nil, errCouldNotParse
	}
	return nil, user.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	for _, result := range f.store.List() {
		if item, ok := result.(user.User); ok {
			if email == item.Email {
				return &item, nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}

	return nil, user.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	for _, result := range f.store.List() {
		if item, ok := result.(user.User); ok {
			if username == item.Username {
				return &item, nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}

	return nil, user.ErrCouldNotFind
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *user.Finder {
	return user.NewFinder(&findRepository{store})
}
