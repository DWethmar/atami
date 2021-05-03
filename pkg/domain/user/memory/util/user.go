package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
)

// AddTestUser adds test user to store with ID = 1
func AddTestUser(store *memstore.Store, ID int) *user.User {
	user := user.User{
		ID:        ID,
		UID:       "UID" + strconv.Itoa(ID),
		Username:  "test",
		Email:     fmt.Sprintf("test%d@test.nl", ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	store.GetUsers().Put(ID, ToMemory(user))
	return &user
}

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
