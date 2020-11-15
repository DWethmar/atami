package memory

import (
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// Find get multiple messages
func (f findRepository) Find() ([]*user.User, error) {
	results := f.store.List()
	items := make([]*user.User, len(results))

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
func (f findRepository) FindByID(ID int) (*user.User, error) {
	if result, ok := f.store.Get(strconv.Itoa(ID)); ok {
		if record, ok := result.(userRecord); ok {
			return recordToUser(record), nil
		}
		return nil, errCouldNotParse
	}
	return nil, user.ErrCouldNotFind
}

// FindByID get one message
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return UID == record.UID
	})
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return email == record.Email
	})
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	return filterList(f.store.List(), func(record userRecord) bool {
		return username == record.Username
	})
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *user.Finder {
	return user.NewFinder(&findRepository{store})
}
