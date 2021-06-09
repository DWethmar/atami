package user

import (
	"fmt"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

// NewUserFixture create new message fixture
func NewUserFixture(ID entity.ID) *User {
	return &User{
		ID:        ID,
		UID:       entity.NewUID(),
		Username:  fmt.Sprintf("user%d", ID),
		Email:     fmt.Sprintf("user%d", ID),
		Password:  "abdefABCDEF1234!@#$",
		Biography: "biography text",
		CreatedAt: entity.Now(),
		UpdatedAt: entity.Now(),
	}
}
