package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/memory/util"
	"github.com/dwethmar/atami/pkg/memstore"
)

// FindUser finds user and parses it to message.User
func FindUser(store *memstore.UserStore, userID int) (*message.User, error) {
	// find and set user
	if r, ok := store.Get(userID); ok {
		user := util.FromMemory(r)
		return ToMsgUser(user), nil
	}
	return nil, fmt.Errorf("Could not find user with ID %d in memory store", userID)
}

// AddTestUser adds test user to store with ID = 1
func AddTestUser(store *memstore.Store, ID int) {
	store.GetUsers().Put(ID, util.ToMemory(user.User{
		ID:        ID,
		UID:       "UID" + strconv.Itoa(ID),
		Username:  "test",
		Email:     fmt.Sprintf("test_%v@test.nl", ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}))
}
