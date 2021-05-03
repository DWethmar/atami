package test

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

// TestFindByID tests the find by id function.
func TestFindByID(t *testing.T, finder *user.Finder, ID int) {
	m, err := finder.FindByID(ID)
	assert.NoError(t, err)
	assert.Equal(t, ID, m.ID)
}

// TestFindByUID tests the find by UID function.
func TestFindByUID(t *testing.T, finder *user.Finder, UID string) {
	m, err := finder.FindByUID(UID)
	assert.NoError(t, err)
	assert.Equal(t, UID, m.UID)
}

// TestUserNotFound tests the ReadOne function for a not found error.
func TestUserNotFound(t *testing.T, finder *user.Finder) {
	_, err := finder.FindByID(0)
	assert.Equal(t, user.ErrCouldNotFind, err)
}

// TestFind tests the ReadOne function.
func TestFind(t *testing.T, finder *user.Finder, length int, users []user.User) {
	list, err := finder.Find()

	assert.Nil(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, user := range list {
			assert.True(t, users[i].Equal(*user), "users are not equal")
		}
	}
}

// TestFindByEmail tests the search function.
func TestFindByEmail(t *testing.T, finder *user.Finder, email string) {
	result, err := finder.FindByEmail(email)
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, email, result.Email)
		assert.Empty(t, result.Password)
	}
}

// TestFindByEmailWithPassword tests the search function.
func TestFindByEmailWithPassword(t *testing.T, finder *user.Finder, email string) {
	result, err := finder.FindByEmailWithPassword(email)
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, email, result.Email)
		assert.NotEmpty(t, result.Password)
	}
}

// TestFindByUsername tests the search function.
func TestFindByUsername(t *testing.T, finder *user.Finder, username string) {
	result, err := finder.FindByUsername(username)
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, username, result.Username)
		assert.Empty(t, result.Password)
	}
}
