package memory

import (
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/memory/util"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// Find get multiple messages
func (f findRepository) Find() ([]*user.User, error) {
	results := f.store.GetUsers().All()
	items := make([]*user.User, len(results))

	for i, r := range f.store.GetUsers().All() {
		user := util.FromMemory(r)
		items[i] = &user
	}

	return items, nil
}

// FindByID get one message
func (f findRepository) FindByID(ID int) (*user.User, error) {
	users := f.store.GetUsers()

	if r, ok := users.Get(ID); ok {
		user := util.FromMemory(r)
		return &user, nil
	}

	return nil, user.ErrCouldNotFind
}

// FindByID get one message
func (f findRepository) FindByUID(UID string) (*user.User, error) {
	u, err := filterList(f.store.GetUsers().All(), func(record user.User) bool {
		return UID == record.UID
	})
	if err == nil && u != nil {
		u.Password = ""
		return u, nil
	}
	return nil, err
}

// FindByEmail func
func (f *findRepository) FindByEmail(email string) (*user.User, error) {
	u, err := filterList(f.store.GetUsers().All(), func(record user.User) bool {
		return email == record.Email
	})
	if err == nil && u != nil {
		u.Password = ""
		return u, nil
	}
	return nil, err
}

// FindByEmailWithPassword func
func (f *findRepository) FindByEmailWithPassword(email string) (*user.User, error) {
	u, err := filterList(f.store.GetUsers().All(), func(record user.User) bool {
		return email == record.Email
	})
	if err == nil && u != nil {
		return u, nil
	}
	return nil, err
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	u, err := filterList(f.store.GetUsers().All(), func(record user.User) bool {
		return username == record.Username
	})
	if err == nil && u != nil {
		u.Password = ""
		return u, nil
	}
	return nil, err
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Store) *user.Finder {
	return user.NewFinder(&findRepository{store})
}
