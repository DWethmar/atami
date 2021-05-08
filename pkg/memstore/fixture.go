package memstore

import (
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

// NewUserFixture create new user
func NewUserFixture(ID entity.ID) *User {
	return &User{
		ID:        ID,
		UID:       entity.NewUID(),
		Username:  fmt.Sprintf("user%d", ID),
		Password:  "abcABC123!@#",
		Email:     "fixtureuser@test.nl",
		Biography: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// NewMessageFixture create new user
func NewMessageFixture(ID entity.ID, createdByUserID entity.ID) *Message {
	return &Message{
		ID:              ID,
		UID:             entity.NewUID(),
		Text:            fmt.Sprintf("text %d", ID),
		CreatedByUserID: createdByUserID,
		CreatedAt:       time.Now(),
	}
}
