package message

import (
	"github.com/dwethmar/atami/pkg/memstore"
)

// MessageToMemoryMap maps a message to memory
func messageToMemoryMap(m Message) *memstore.Message {
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
func messageFromMemoryMap(m memstore.Message) *Message {
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
