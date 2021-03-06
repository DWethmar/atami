package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/dwethmar/atami/pkg/memstore"
)

func generateTestUsers(size int) []user.CreateUser {
	users := make([]user.CreateUser, size)
	for i := 0; i < size; i++ {
		users[i] = user.CreateUser{
			Username: fmt.Sprintf("username_%d", i+1),
			Email:    fmt.Sprintf("test-%d@test.com", i+1),
			Password: fmt.Sprintf("abcdef1234!@#$ABCD-%d", i+1),
		}
	}
	return users
}

func setup() (*user.Finder, []user.User) {
	memstore := memstore.NewStore()
	creator := NewCreator(memstore, NewFinder(memstore))

	users := make([]user.User, 100)
	for i, testUser := range generateTestUsers(100) {
		user, err := creator.Create(testUser)
		if err != nil {
			fmt.Printf("error: %s", err)
			panic(1)
		}
		users[i] = *user
	}
	return NewFinder(memstore), users
}

func TestFind(t *testing.T) {
	finder, users := setup()
	test.TestFind(t, finder, 100, users)
}

func TestFindByUID(t *testing.T) {
	finder, users := setup()
	test.TestFindByUID(t, finder, users[0].UID)
}

func TestFindByID(t *testing.T) {
	finder, users := setup()
	test.TestFindByID(t, finder, users[0].ID)
}

func TestUserNotFound(t *testing.T) {
	finder, _ := setup()
	test.TestUserNotFound(t, finder)
}

func TestFindByEmail(t *testing.T) {
	finder, _ := setup()
	test.TestFindByEmail(t, finder, "test-44@test.com")
}

func TestFindByEmailWithPassword(t *testing.T) {
	finder, _ := setup()
	test.TestFindByEmailWithPassword(t, finder, "test-44@test.com")
}

func TestFindByUsername(t *testing.T) {
	finder, _ := setup()
	test.TestFindByUsername(t, finder, "username_44")
}
