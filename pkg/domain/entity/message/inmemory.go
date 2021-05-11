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
}

//NewInMemoryRepo create new repository
func NewInMemoryRepo(memStore *memstore.Memstore) Repository {
	return &inMemoryRepo{
		memStore: memStore,
	}
}

func (r *inMemoryRepo) Get(ID entity.ID) (*Message, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()

	if r, ok := messages.Get(ID); ok {
		msg := messageFromMemoryMap(&r)

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

	var low = offset
	var high = offset + limit

	l := messages.Len()

	if l == 0 {
		return []*Message{}, nil
	}

	if low > uint(l) {
		return []*Message{}, nil
	}

	if high > uint(l) {
		high = uint(l)
	}

	all, _ := messages.All()
	sort.Slice(all, func(i, j int) bool {
		var a = all[i]
		var b = all[j]
		return a.ID > b.ID
	})

	items := make([]*Message, 0)
	for _, r := range all[low:high] {
		msg := messageFromMemoryMap(&r)
		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.CreatedBy = *user
		} else {
			return nil, err
		}
		items = append(items, msg)
	}

	return items, nil
}

func (r *inMemoryRepo) Update(message *Message) error {
	messages := r.memStore.GetMessages()
	if _, ok := messages.Get(message.ID); !ok {
		return domain.ErrNotFound
	}

	mapped := messageToMemoryMap(message)
	if messages.Delete(message.ID) && !messages.Put(message.ID, *mapped) {
		return domain.ErrCannotBeUpdated
	}
	return nil
}

func (r *inMemoryRepo) Create(message *Message) (entity.ID, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()

	if _, ok := users.Get(message.CreatedByUserID); !ok {
		return 0, errors.New("user not found :(")
	}

	message.ID = messages.Len() + 1
	mapped := messageToMemoryMap(message)
	messages.Put(message.ID, *mapped)

	if r, ok := messages.Get(message.ID); ok {
		msg := messageFromMemoryMap(&r)
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
		message := messageFromMemoryMap(&item)
		if filterFn(message) {
			return message, nil
		}
	}
	return nil, domain.ErrNotFound
}

// MessageToMemoryMap maps a message to memory
func messageToMemoryMap(m *Message) *memstore.Message {
	return &memstore.Message{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

// MessageFromMemoryMap maps a message from memory
func messageFromMemoryMap(m *memstore.Message) *Message {
	return &Message{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

// UserFromMemoryMap maps a message from memory
func userToMemoryMap(m User) *memstore.User {
	return &memstore.User{
		ID:       m.ID,
		UID:      m.UID,
		Username: m.Username,
	}
}

// UserFromMemoryMap maps a message from memory
func userFromMemoryMap(m memstore.User) *User {
	return &User{
		ID:       m.ID,
		UID:      m.UID,
		Username: m.Username,
	}
}
