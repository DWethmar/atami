package memory

import (
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

// findRepository reads messages from memory
type findRepository struct {
	store *memstore.Memstore
}

// Find get multiple messages
func (f findRepository) Find() ([]*user.User, error) {
	users, err := f.store.GetUsers().All()
	if err != nil {
		return nil, err
	}

	items := make([]*user.User, len(users))

	for i, r := range users {
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
	users, err := f.store.GetUsers().All()
	if err != nil {
		return nil, err
	}

	u, err := filterList(users, func(record user.User) bool {
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
	users, err := f.store.GetUsers().All()
	if err != nil {
		return nil, err
	}

	u, err := filterList(users, func(record user.User) bool {
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
	users, err := f.store.GetUsers().All()
	if err != nil {
		return nil, err
	}

	u, err := filterList(users, func(record user.User) bool {
		return email == record.Email
	})
	if err == nil && u != nil {
		return u, nil
	}
	return nil, err
}

// FindByEmail func
func (f *findRepository) FindByUsername(username string) (*user.User, error) {
	users, err := f.store.GetUsers().All()
	if err != nil {
		return nil, err
	}

	u, err := filterList(users, func(record user.User) bool {
		return username == record.Username
	})
	if err == nil && u != nil {
		u.Password = ""
		return u, nil
	}
	return nil, err
}

// NewFinder return a new in memory listin repository
func NewFinder(store *memstore.Memstore) *user.Finder {
	return user.NewFinder(&findRepository{store})
}
