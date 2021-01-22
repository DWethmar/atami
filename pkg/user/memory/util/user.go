package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
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
	store.GetUsers().Put(strconv.Itoa(ID), user)
	return &user
}
