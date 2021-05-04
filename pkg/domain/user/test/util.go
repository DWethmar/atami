package test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/domain/user"
)

// AddTestUser adds test user to store with ID = 1
func AddTestUser(creator *user.Creator, ID int) *user.User {
	u, _ := creator.Create(user.CreateUser{
		UID:       "UID" + strconv.Itoa(ID),
		Username:  "test",
		Email:     fmt.Sprintf("test%d@test.nl", ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return u
}
