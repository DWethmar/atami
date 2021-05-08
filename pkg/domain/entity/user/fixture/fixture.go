package fixture

import (
	"fmt"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/domain/entity/user"
)

// NewUserFixture create new message fixture
func NewUserFixture(ID entity.ID) *user.User {
	return &user.User{
		ID:        ID,
		UID:       entity.NewUID(),
		Username:  fmt.Sprintf("user%d", ID),
		Email:     fmt.Sprintf("user%d", ID),
		Password:  "abdefABCDEF1234!@#$",
		Biography: "biography text",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
