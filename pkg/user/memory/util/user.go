package util

import (
	"strconv"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

// AddTestUser adds test user to store with ID = 1
func AddTestUser(store *memstore.Store, ID int) {
	store.GetUsers().Put(strconv.Itoa(ID), user.User{
		ID:        ID,
		UID:       "UID" + strconv.Itoa(ID),
		Username:  "test",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
