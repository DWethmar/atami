package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadOne tests the ReadOne function.
func TestReadOne(t *testing.T, reader *Reader, ID ID, messages Message) {
	m, err := reader.ReadOne(ID)
	assert.Nil(t, err)

	assert.Equal(t, messages.ID, m.ID)
	assert.Equal(t, messages.Content, m.Content)
	assert.Equal(t, messages.CreatedAt, m.CreatedAt)
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, reader *Reader) {
	_, err := reader.ReadOne(ID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestReadAll tests the ReadOne function.
func TestReadAll(t *testing.T, reader *Reader, length int, messages []Message) {
	list, err := reader.ReadAll()

	assert.Nil(t, err)
	assert.Equal(t, length, len(list))

	for i, message := range list {
		assert.Equal(t, messages[i].ID, message.ID)
		assert.Equal(t, messages[i].UID, message.UID)
		assert.Equal(t, messages[i].Content, message.Content)
	}
}
