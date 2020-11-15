package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/stretchr/testify/assert"
)

func generateTestUsers(size int) []auth.CreateUser {
	users := make([]auth.CreateUser, size)
	for i := 0; i < size; i++ {
		users[i] = auth.CreateUser{
			Username: fmt.Sprintf("username_%d", i+1),
			Email:    fmt.Sprintf("test-%d@test.com", i+1),
			Password: fmt.Sprintf("abcdef1234!@#$ABCD-%d", i+1),
		}
	}
	return users
}

func setup(db *sql.DB) (*auth.Finder, []auth.User) {
	service := New(db)
	users := make([]auth.User, 100)
	for i, testUser := range generateTestUsers(100) {
		user, err := service.Register(testUser)
		if err != nil {
			fmt.Printf("error: %s", err)
			panic(1)
		}
		users[i] = *user
	}
	return NewFinder(db), users
}

func TestFind(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		var existingUsers []auth.User
		finder := NewFinder(db)
		if users, err := finder.Find(); err == nil {
			for _, u := range users {
				existingUsers = append(existingUsers, *u)
			}
		} else {
			assert.Fail(t, "could not query users")
		}

		finder, users := setup(db)
		users = append(existingUsers, users...)
		auth.TestFind(t, finder, 100+len(existingUsers), users)

		return nil
	}))
}

func TestFindByID(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, users := setup(db)
		auth.TestFindByID(t, finder, users[0].ID)

		return nil
	}))
}

func TestUserNotFound(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, _ := setup(db)
		auth.TestUserNotFound(t, finder)

		return nil
	}))
}

func TestFindByEmail(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, _ := setup(db)
		auth.TestFindByEmail(t, finder, "test-44@test.com")

		return nil
	}))
}

func TestFindByUsername(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, _ := setup(db)
		auth.TestFindByUsername(t, finder, "username_44")

		return nil
	}))
}
