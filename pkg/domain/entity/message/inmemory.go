package message

import (
	"errors"
	"fmt"

	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/memstore"
)

type inMemoryRepo struct {
	memStore *memstore.Memstore
	newID    entity.ID
}

//NewinMemoryRepoRepository create new repository
func NewinMemoryRepoRepository(memStore *memstore.Memstore) Repository {
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
			msg.User = user
		} else {
			return nil, err
		}

		return &msg, nil
	}

	return nil, domain.ErrNotFound
}

func (r *inMemoryRepo) GetByUID(UID entity.UID) (*Message, error) {
	messages, err := r.memStore.GetMessages().All()
	if err != nil {
		return nil, err
	}

	msg, err := filterMessagesFromMemory(messages, func(message Message) bool {
		return UID == message.UID
	})

	if msg == nil {
		return nil, err
	}

	users := r.memStore.GetUsers()

	if err == nil {
		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.User = user
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

	paged, err := messages.Slice(offset, limit)
	if err != nil {
		return nil, err
	}

	items := make([]*Message, len(paged))

	for i, r := range paged {
		msg := messageFromMemoryMap(r)

		if user, err := findUserInMemstore(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}

		items[i] = &msg
	}

	return items, nil
}

func (r *inMemoryRepo) Update(ID entity.ID, change Update) error {
	messages := r.memStore.GetMessages()
	message, err := r.Get(ID)
	if err != nil {
		return err
	}
	message.Apply(change)
	mapped := messageToMemoryMap(*message)
	if messages.Delete(message.ID) && !messages.Put(message.ID, mapped) {
		return errors.New("Could not update message")
	}
	return nil
}

func (r *inMemoryRepo) Create(create Create) (entity.ID, error) {
	messages := r.memStore.GetMessages()
	users := r.memStore.GetUsers()
	_, ok := users.Get(create.CreatedByUserID)

	r.newID++

	if !ok {
		return 0, errors.New("user not found")
	}

	msg := Message{
		ID:              r.newID,
		UID:             create.UID,
		Text:            create.Text,
		CreatedByUserID: create.CreatedByUserID,
		CreatedAt:       create.CreatedAt,
	}

	mapped := messageToMemoryMap(msg)
	messages.Put(msg.ID, mapped)

	if r, ok := messages.Get(msg.ID); ok {
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
func findUserInMemstore(store *memstore.UserStore, userID entity.ID) (User, error) {
	if r, ok := store.Get(userID); ok {
		user := userFromMemoryMap(r)
		return user, nil
	}
	return User{}, fmt.Errorf("Could not find user with ID %d in memory store", userID)
}

func filterMessagesFromMemory(list []memstore.Message, filterFn func(Message) bool) (*Message, error) {
	for _, item := range list {
		message := messageFromMemoryMap(item)
		if filterFn(message) {
			return &message, nil
		}
	}
	return nil, domain.ErrNotFound
}
