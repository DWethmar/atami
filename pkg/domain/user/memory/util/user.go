package util

import (
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
)

// ToMemory maps a message to memory
func ToMemory(m user.User) memstore.User {
	return memstore.User{
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

// FromMemory maps a message from memory
func FromMemory(m memstore.User) user.User {
	return user.User{
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
