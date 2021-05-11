package user

import (
	"sort"

	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
)

type inMemoryRepo struct {
	memStore *memstore.Memstore
}

//NewInMemoryRepo create new repository
func NewInMemoryRepo(memStore *memstore.Memstore) Repository {
	return &inMemoryRepo{
		memStore: memStore,
	}
}

func (r *inMemoryRepo) Get(ID entity.ID) (*User, error) {
	users := r.memStore.GetUsers()

	if r, ok := users.Get(ID); ok {
		user := fromMemory(r)
		return user, nil
	}

	return nil, domain.ErrNotFound
}

func (r *inMemoryRepo) GetByUID(UID entity.UID) (*User, error) {
	users, err := r.memStore.GetUsers().All()
	if err != nil {
		return nil, err
	}

	u, err := filterList(users, func(record *User) bool {
		return UID == record.UID
	})

	if err == nil && u != nil {
		u.Password = ""
		return u, nil
	}
	return nil, err
}

func (r *inMemoryRepo) List(limit, offset uint) ([]*User, error) {
	users := r.memStore.GetUsers()

	var low = offset
	var high = offset + limit

	l := users.Len()

	if l == 0 {
		return []*User{}, nil
	}

	if low > uint(l) {
		return []*User{}, nil
	}

	if high > uint(l) {
		high = uint(l)
	}

	all, _ := users.All()
	sort.Slice(all, func(i, j int) bool {
		var a = all[i]
		var b = all[j]
		return a.ID > b.ID
	})

	items := make([]*User, 0)
	for _, r := range all[low:high] {
		items = append(items, fromMemory(r))
	}

	return items, nil
}

func (r *inMemoryRepo) Update(user *User) error {
	userStore := r.memStore.GetUsers()
	if _, ok := userStore.Get(user.ID); !ok {
		return domain.ErrNotFound
	}
	users, err := userStore.All()
	if err != nil {
		return nil
	}

	// Check for user with same username or email
	if match, err := filterList(users, func(record *User) bool {
		return user.Username == record.Username || user.Email == record.Email
	}); match != nil {
		if match.Email == user.Email {
			return ErrEmailAlreadyTaken
		}
		if match.Username == user.Username {
			return ErrUsernameAlreadyTaken
		}
		return err
	} else if err != nil && err != domain.ErrNotFound {
		return err
	}

	mapped := toMemory(user)
	if userStore.Delete(mapped.ID) && !userStore.Put(mapped.ID, *mapped) {
		return domain.ErrCannotBeUpdated
	}
	return nil
}

func (r *inMemoryRepo) Create(user *User) (entity.ID, error) {
	userStore := r.memStore.GetUsers()
	users, err := userStore.All()
	if err != nil {
		return 0, err
	}

	// Check for user with same username or email
	if match, err := filterList(users, func(record *User) bool {
		return user.Username == record.Username || user.Email == record.Email
	}); match != nil && err == nil {
		if match.Email == user.Email {
			return 0, ErrEmailAlreadyTaken
		}
		if match.Username == user.Username {
			return 0, ErrUsernameAlreadyTaken
		}
		return 0, err
	} else if err != nil && err != domain.ErrNotFound {
		return 0, err
	}

	user.ID = userStore.Len() + 1
	userStore.Put(user.ID, *toMemory(user))

	return user.ID, nil
}

func (r *inMemoryRepo) Delete(ID entity.ID) error {
	userStore := r.memStore.GetUsers()
	if userStore.Delete(ID) {
		return nil
	}
	return domain.ErrCannotBeDeleted
}

// toMemory maps a message to memory
func toMemory(m *User) *memstore.User {
	return &memstore.User{
		ID:        m.ID,
		UID:       m.UID,
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		Biography: m.Biography,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// fromMemory maps a message from memory
func fromMemory(m memstore.User) *User {
	return &User{
		ID:        m.ID,
		UID:       m.UID,
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		Biography: m.Biography,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func filterList(list []memstore.User, filterFunc func(*User) bool) (*User, error) {
	for _, item := range list {
		user := fromMemory(item)
		if filterFunc(user) {
			return user, nil
		}
	}
	return nil, domain.ErrNotFound
}
