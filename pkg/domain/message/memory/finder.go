package memory

import (
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

// readerRepository reads messages from memory
type findRepository struct {
	store *memstore.Memstore
}

// FindByID get one message
func (i *findRepository) FindByUID(UID string) (*message.Message, error) {
	messages, err := i.store.GetMessages().All()
	if err != nil {
		return nil, err
	}

	msg, err := filterList(messages, func(record message.Message) bool {
		return UID == record.UID
	})

	if msg == nil {
		return nil, err
	}

	users := i.store.GetUsers()

	if err == nil {
		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return msg, nil
}

// FindByID get one message
func (i *findRepository) FindByID(ID int) (*message.Message, error) {
	messages := i.store.GetMessages()
	users := i.store.GetUsers()

	if r, ok := messages.Get(ID); ok {
		msg := util.FromMemory(r)
		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}

		return &msg, nil
	}
	return nil, message.ErrCouldNotFind
}

// FindAll get multiple messages
func (i *findRepository) Find(limit, offset int) ([]*message.Message, error) {
	messages := i.store.GetMessages()
	users := i.store.GetUsers()

	if len := messages.Len(); len == 0 {
		return nil, nil
	} else if offset+limit > len {
		limit = len - offset
	}

	paged, err := messages.Slice(offset, limit)
	if err != nil {
		return nil, err
	}

	items := make([]*message.Message, len(paged))

	for i, r := range paged {

		msg := util.FromMemory(r)

		// fmt.Println("------------------------------------------------------------------------")
		// fmt.Println(fmt.Sprintf("Created BY %d", msg.CreatedByUserID))
		// fmt.Println(util.FindUser(users, msg.CreatedByUserID))
		// test, _ := users.Get(msg.CreatedByUserID)
		// fmt.Println(test)
		// fmt.Println("------------------------------------------------------------------------")

		if user, err := util.FindUser(users, msg.CreatedByUserID); err == nil {
			msg.User = user
		} else {
			return nil, err
		}
		items[i] = &msg
	}

	return items, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(store *memstore.Memstore) *message.Finder {
	return message.NewFinder(&findRepository{store})
}
