package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadOne tests the ReadOne function.
func TestReadOne(t *testing.T, finder *Finder, ID ID, messages Message) {
	m, err := finder.FindByID(ID)
	assert.Nil(t, err)

	assert.Equal(t, messages.ID, m.ID)
	assert.Equal(t, messages.Text, m.Text)
	assert.Equal(t, messages.CreatedAt, m.CreatedAt)
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(ID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestReadAll tests the ReadOne function.
func TestReadAll(t *testing.T, finder *Finder, length int, messages []Message) {
	list, err := finder.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, length, len(list))

	for i, message := range list {
		assert.Equal(t, messages[i].ID, message.ID)
		assert.Equal(t, messages[i].UID, message.UID)
		assert.Equal(t, messages[i].Text, message.Text)
	}
}
