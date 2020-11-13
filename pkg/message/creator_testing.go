package message

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, creator *Creator, newMessage NewMessage) {
	message, err := creator.Create(newMessage)

	assert.Nil(t, err)
	assert.Equal(t, message.ID, model.MessageID(1))
	assert.Equal(t, message.Text, newMessage.Content)
	time.Sleep(1)
	assert.True(t, time.Now().After(message.CreatedAt))
}
