package message

import (
	"errors"
	"fmt"
	"sort"

	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
)

type inMemoryRepo struct {
	memStore *memstore.Memstore
	newID    entity.ID
}

//NewInMemoryRepo create new repository
func NewInMemoryRepo(memStore *memstore.Memstore) Repository {
	return &inMemoryRepo{
		memStore: memStore,
		newID:    0,
	}
}

func (r *inMemoryRepo) Get(ID entity.ID) (*Message, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()

	if r, ok := messages.Get(ID); ok {
		msg := messageFromMemoryMap(r)

		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.CreatedBy = *user
		} else {
			return nil, err
		}

		return msg, nil
	}

	return nil, domain.ErrNotFound
}

func (r *inMemoryRepo) GetByUID(UID entity.UID) (*Message, error) {
	messages, err := r.memStore.GetMessages().All()
	if err != nil {
		return nil, err
	}

	msg, err := filterMessagesFromMemory(messages, func(message *Message) bool {
		return UID == message.UID
	})

	if msg == nil {
		return nil, err
	}

	users := r.memStore.GetUsers()

	if err == nil {
		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.CreatedBy = *user
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return msg, nil
}

func (r *inMemoryRepo) List(limit, offset uint) ([]*Message, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()

	if len := messages.Len(); len == 0 {
		return []*Message{}, nil
	} else if offset > uint(len) {
		return []*Message{}, nil
	} else if offset+limit > uint(len) {
		limit = uint(len)
	}

	paged, err := messages.Slice(offset, offset + limit)
	if err != nil {
		return nil, err
	}

	items := make([]*Message, len(paged))
	for i, r := range paged {
		msg := messageFromMemoryMap(r)
		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.CreatedBy = *user
		} else {
			return nil, err
		}
		items[i] = msg
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ID > items[j].ID
	})

	return items, nil
}

func (r *inMemoryRepo) Update(message *Message) error {
	messages := r.memStore.GetMessages()
	mapped := messageToMemoryMap(*message)
	if messages.Delete(message.ID) && !messages.Put(message.ID, *mapped) {
		return errors.New("could not update message")
	}
	return nil
}

func (r *inMemoryRepo) Create(message *Message) (entity.ID, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()

	if _, ok := users.Get(message.CreatedByUserID); !ok {
		return 0, errors.New("user not found")
	}

	r.newID++
	message.ID = r.newID

	mapped := messageToMemoryMap(*message)
	messages.Put(message.ID, *mapped)

	if r, ok := messages.Get(message.ID); ok {
		msg := messageFromMemoryMap(r)
		return msg.ID, nil
	}

	return 0, errors.New("could not find message")
}

func (r *inMemoryRepo) Delete(ID entity.ID) error {
	messages := r.memStore.GetMessages()
	if messages.Delete(ID) {
		return nil
	}
	return domain.ErrCannotBeDeleted
}

// findUserInMemstore finds user and parses it to User
func findUserInMemstore(store *memstore.UserStore, userID entity.ID) (*User, error) {
	if r, ok := store.Get(userID); ok {
		user := userFromMemoryMap(r)
		return user, nil
	}
	return nil, fmt.Errorf("could not find user with ID %d in memory store", userID)
}

func filterMessagesFromMemory(list []memstore.Message, filterFn func(*Message) bool) (*Message, error) {
	for _, item := range list {
		message := messageFromMemoryMap(item)
		if filterFn(message) {
			return message, nil
		}
	}
	return nil, domain.ErrNotFound
}
