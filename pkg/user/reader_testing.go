package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadOne tests the ReadOne function.
func TestReadOne(t *testing.T, reader *Reader, ID ID, user User) {
	m, err := reader.ReadOne(ID)
	assert.Nil(t, err)

	assert.Equal(t, user.ID, m.ID)
	assert.Equal(t, user.Email, m.Email)
	assert.Equal(t, user.CreatedAt, m.CreatedAt)
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, reader *Reader) {
	_, err := reader.ReadOne(ID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestReadAll tests the ReadOne function.
func TestReadAll(t *testing.T, reader *Reader, length int, users []User) {
	list, err := reader.ReadAll()

	assert.Nil(t, err)
	assert.Equal(t, length, len(list))

	for i, user := range list {
		assert.Equal(t, users[i].ID, user.ID)
		assert.Equal(t, users[i].UID, user.UID)
		assert.Equal(t, users[i].Email, user.Email)
	}
}
