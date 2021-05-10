package message

import (
	"github.com/dwethmar/atami/pkg/domain/entity"
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

func insertRowMap(row Row) (entity.ID, error) {
	var ID entity.ID
	if err := row.Scan(
		&ID,
	); err != nil {
		return 0, err
	}
	return ID, nil
}

func messageWithUserRowMap(row Row) (*Message, error) {
	e := &Message{
		CreatedBy: User{},
	}
	if err := row.Scan(
		&e.ID,
		&e.UID,
		&e.Text,
		&e.CreatedByUserID,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.CreatedBy.ID,
		&e.CreatedBy.UID,
		&e.CreatedBy.Username,
	); err != nil {
		return nil, err
	}
	e.CreatedAt.UTC()
	e.UpdatedAt.UTC()
	return e, nil
}
