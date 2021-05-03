package util

import (
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
)

// ToMsgUser from the memstore to message user
func ToMsgUser(user user.User) *message.User {
	return &message.User{
		ID:       user.ID,
		UID:      user.UID,
		Username: user.Username,
	}
}

// ToMemory maps a message to memory
func ToMemory(m message.Message) memstore.Message {
	return memstore.Message{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
	}
}

// FromMemory maps a message from memory
func FromMemory(m memstore.Message) message.Message {
	return message.Message{
		ID:              m.ID,
		UID:             m.UID,
		Text:            m.Text,
		CreatedByUserID: m.CreatedByUserID,
		CreatedAt:       m.CreatedAt,
	}
}
