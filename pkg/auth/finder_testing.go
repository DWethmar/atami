package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindByID tests the find by id function.
func TestFindByID(t *testing.T, finder *Finder, ID ID) {
	m, err := finder.FindByID(ID)
	assert.Nil(t, err)
	assert.Equal(t, ID, m.ID)
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(ID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestFindAll tests the ReadOne function.
func TestFindAll(t *testing.T, finder *Finder, length int, users []User) {
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
	assert.Nil(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, email, result.Email)
	}
}

// TestFindByUsername tests the search function.
func TestFindByUsername(t *testing.T, finder *Finder, username string) {
	result, err := finder.FindByUsername(username)
	assert.Nil(t, err)
	if assert.NotNil(t, result) {
		assert.Equal(t, username, result.Username)
	}
}
