package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadOne tests the ReadOne function.
func TestReadOne(t *testing.T, repo ReaderRepository, ID ID, messages Message) {
	reader := NewReader(repo)
	m, err := reader.ReadOne(ID)
	assert.Nil(t, err)

	assert.Equal(t, messages.ID, m.ID)
	assert.Equal(t, messages.Content, m.Content)
	assert.Equal(t, messages.CreatedAt, m.CreatedAt)
}

// TestReadAll tests the ReadOne function.
func TestReadAll(t *testing.T, repo ReaderRepository, length int, messages []Message) {
	reader := NewReader(repo)
	list, err := reader.ReadAll()

	assert.Nil(t, err)
	assert.Equal(t, length, len(list))

	for i, message := range list {
		assert.Equal(t, messages[i].ID, message.ID)
		assert.Equal(t, messages[i].UID, message.UID)
		assert.Equal(t, messages[i].Content, message.Content)
	}
}
