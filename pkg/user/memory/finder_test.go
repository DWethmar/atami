package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func generateTestUsers(size int) []user.CreateUser {
	users := make([]user.CreateUser, size)
	for i := 0; i < size; i++ {
		users[i] = user.CreateUser{
			Username:       fmt.Sprintf("username_%d", i+1),
			Email:          fmt.Sprintf("test-%d@test.com", i+1),
			HashedPassword: fmt.Sprintf("abcdef1234!@#$ABCD-%d", i+1),
		}
	}
	return users
}

func setup() (*user.Finder, []user.User) {
	store := memstore.New()
	service := New(store)
	users := make([]user.User, 100)
	for i, testUser := range generateTestUsers(100) {
		user, err := service.Create(testUser)
		if err != nil {
			fmt.Printf("error: %s", err)
			panic(1)
		}
		users[i] = *user
	}
	return NewFinder(store), users
}

func TestFind(t *testing.T) {
	finder, users := setup()
	user.TestFind(t, finder, 100, users)
}

func TestFindByID(t *testing.T) {
	finder, users := setup()
	user.TestFindByID(t, finder, users[0].ID)
}

func TestUserNotFound(t *testing.T) {
	finder, _ := setup()
	user.TestUserNotFound(t, finder)
}

func TestFindByEmail(t *testing.T) {
	finder, _ := setup()
	user.TestFindByEmail(t, finder, "test-44@test.com")
}

func TestFindByUsername(t *testing.T) {
	finder, _ := setup()
	user.TestFindByUsername(t, finder, "username_44")
}
