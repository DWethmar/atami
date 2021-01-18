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
	results := f.store.GetUsers().All()
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
	users := f.store.GetUsers()

	if result, ok := users.Get(strconv.Itoa(ID)); ok {
		if record, ok := result.(userRecord); ok {
			return recordToUser(record), nil
		}
		return nil, errCouldNotParse
	}
	return nil, user.ErrCouldNotFind
}

// FindByID get one message
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	return filterList(f.store.GetUsers().All(), func(record userRecord) bool {
		return UID == record.UID
	})
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	for _, item := range f.store.GetUsers().All() {
		if record, ok := item.(userRecord); ok {
			if record.Email == email {
				usr := recordToUser(record)
				return usr, nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}
	return nil, user.ErrCouldNotFind
}

// FindByEmailWithPassword func
func (f *findRepository) FindByEmailWithPassword(email string) (*user.User, error) {
	for _, item := range f.store.GetUsers().All() {
		if record, ok := item.(userRecord); ok {
			if record.Email == email {
				usr := recordToUser(record)
				usr.Password = record.Password
				return usr, nil
			}
		} else {
			return nil, errCouldNotParse
		}
	}
	return nil, user.ErrCouldNotFind
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	return filterList(f.store.GetUsers().All(), func(record userRecord) bool {
		return username == record.Username
	})
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *user.Finder {
	return user.NewFinder(&findRepository{store})
}
