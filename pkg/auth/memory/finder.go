package memory

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/model"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

func filterList(list []interface{}, filterFn func(userRecord) bool) (*auth.User, error) {
	for _, item := range list {
		if record, ok := item.(userRecord); ok {
			if filterFn(record) {
				return recordToUser(record), nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}
	return nil, auth.ErrCouldNotFind
}

// Find get multiple messages
func (f findRepository) Find() ([]*auth.User, error) {
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
func (f findRepository) FindByID(ID model.UserID) (*auth.User, error) {
	if result, ok := f.store.Get(ID.String()); ok {
		if record, ok := result.(userRecord); ok {
			return recordToUser(record), nil
		}
		return nil, errCouldNotParse
	}
	return nil, auth.ErrCouldNotFind
}

// FindByID get one message
func (f findRepository) FindByUID(UID model.UserUID) (*auth.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return UID == record.UID
	})
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*auth.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return email == record.Email
	})
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*auth.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return username == record.Username
	})
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *auth.Finder {
	return auth.NewFinder(&findRepository{store})
}
