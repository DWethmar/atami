package memory

import (
	"errors"
	"strconv"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/memory/util"
)

// creatorRepository stores new messages
type creatorRepository struct {
	store *memstore.Store
	newID int
}

// Create new message
func (i *creatorRepository) Create(newMsg message.CreateMessage) (*message.Message, error) {
	messages := i.store.GetMessages()
	users := i.store.GetUsers()

	i.newID++
	msg := message.Message{
		ID:              i.newID,
		UID:             newMsg.UID,
		Text:            newMsg.Text,
		CreatedByUserID: newMsg.CreatedByUserID,
		CreatedAt:       newMsg.CreatedAt,
	}
	idStr := strconv.Itoa(msg.ID)

	if result, ok := users.Get(strconv.Itoa(msg.CreatedByUserID)); ok {
		user, err := util.ToMsgUser(result)
		if err == nil {
			msg.User = user
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("user not found")
	}

	messages.Put(idStr, msg)

	if value, ok := messages.Get(idStr); ok {
		if msg, ok := value.(message.Message); ok {
			return &msg, nil
		}
		return nil, errCouldNotParse
	}

	return nil, errors.New("Could not find message")
}

// NewCreator creates new messages creator.
func NewCreator(store *memstore.Store) *message.Creator {
	return message.NewCreator(&creatorRepository{
		store,
		0,
	})
}
