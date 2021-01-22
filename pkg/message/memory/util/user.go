package util

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user"
)

// FindUser finds user and parses it to message.User
func FindUser(store *memstore.KvStore, userID int) (*message.User, error) {
	// find and set user
	if r, ok := store.Get(strconv.Itoa(userID)); ok {
		if user, err := ToMsgUser(r); err == nil {
			return user, nil
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Could not find user in memory store")
}

// AddTestUser adds test user to store with ID = 1
func AddTestUser(store *memstore.Store, ID int) {
	store.GetUsers().Put(strconv.Itoa(ID), user.User{
		ID:        ID,
		UID:       "UID" + strconv.Itoa(ID),
		Username:  "test",
		Email:     fmt.Sprintf("test_%v@test.nl", ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
