package memory

import (
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/memory/util"
)

// readerRepository reads messages from memory
type findRepository struct {
	store *memstore.Store
}

// FindByID get one message
func (i *findRepository) FindByID(ID int) (*message.Message, error) {
	messages := i.store.GetMessages()
	result, ok := messages.Get(strconv.Itoa(ID))
	if ok {
		if message, ok := result.(message.Message); ok {
			return &message, nil
		}
		return nil, errCouldNotParse
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

	paged := messages.Slice(offset, limit)
	items := make([]*message.Message, len(paged))

	for i, l := range paged {
		if msg, ok := l.(message.Message); ok {

			// find and set user
			if r, ok := users.Get(strconv.Itoa(msg.CreatedByUserID)); ok {
				if user, err := util.ToMsgUser(r); err == nil {
					msg.User = user
				}
			}

			items[i] = &msg
		} else {
			return nil, errCouldNotParse
		}
	}

	return items, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(store *memstore.Store) *message.Finder {
	return message.NewFinder(&findRepository{store})
}
