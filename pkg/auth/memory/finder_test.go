package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

func generateTestUsers(size int) []auth.RegisterUser {
	users := make([]auth.RegisterUser, size)
	for i := 0; i < size; i++ {
		users[i] = auth.RegisterUser{
			Username:      fmt.Sprintf("username_%d", i+1),
			Email:         fmt.Sprintf("test-%d@test.com", i+1),
			PlainPassword: fmt.Sprintf("abcdef1234!@#$ABCD-%d", i+1),
		}
	}
	return users
}

func setup() (*auth.Finder, []auth.User) {
	store := memstore.New()
	service := NewService(store)
	users := make([]auth.User, 100)
	for i, testUser := range generateTestUsers(100) {
		user, err := service.Register(testUser)
		if err != nil {
			fmt.Printf("error: %s", err)
			panic(1)
		}
		users[i] = *user
	}
	return NewFinder(store), users
}

func TestFindAll(t *testing.T) {
	finder, users := setup()
	auth.TestFindAll(t, finder, 100, users)
}

func TestFindByID(t *testing.T) {
	finder, users := setup()
	auth.TestFindByID(t, finder, users[0].ID)
}

func TestUserNotFound(t *testing.T) {
	finder, _ := setup()
	auth.TestUserNotFound(t, finder)
}

func TestFindByEmail(t *testing.T) {
	finder, _ := setup()
	auth.TestFindByEmail(t, finder, "test-44@test.com")
}

func TestFindByUsername(t *testing.T) {
	finder, _ := setup()
	auth.TestFindByUsername(t, finder, "username_44")
}
