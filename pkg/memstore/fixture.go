package memstore

import (
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

// NewFixtureUser create new user
func NewFixtureUser(ID entity.ID) *User {
	return &User{
		ID:        ID,
		UID:       entity.NewUID(),
		Username:  "fixtureuser",
		Password:  "abcABC123!@#",
		Email:     "fixtureuser@test.nl",
		Biography: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
