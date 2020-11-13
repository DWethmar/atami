package auth

import (
	"testing"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TestFindByID tests the find by id function.
func TestFindByID(t *testing.T, finder *Finder, ID model.UserID) {
	m, err := finder.FindByID(ID)
	assert.NoError(t, err)
	assert.Equal(t, ID, m.ID)
}

// TestUserNotFound tests the ReadOne function for a not found error.
func TestUserNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(model.UserID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestFindAll tests the ReadOne function.
func TestFindAll(t *testing.T, finder *Finder, length int, users []model.User) {
	list, err := finder.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, length, len(list))

	for i, user := range list {
		assert.Equal(t, users[i].ID, user.ID)
		assert.Equal(t, users[i].UID, user.UID)
		assert.Equal(t, users[i].Email, user.Email)
	}
}

// TestFindByEmail tests the search function.
func TestFindByEmail(t *testing.T, finder *Finder, email string) {
	result, err := finder.FindByEmail(email)
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, email, result.Email)
	}
}

// TestFindByUsername tests the search function.
func TestFindByUsername(t *testing.T, finder *Finder, username string) {
	result, err := finder.FindByUsername(username)
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, username, result.Username)
	}
}
