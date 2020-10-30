package message

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, repo CreatorRepository, newMessage NewMessage) {
	creator := NewCreator(repo)
	message, err := creator.Create(newMessage)
	assert.Nil(t, err)
	assert.Equal(t, message.ID, ID(1))
	assert.Equal(t, message.Content, newMessage.Content)
	time.Sleep(1)
	assert.True(t, time.Now().After(message.CreatedAt))
}
