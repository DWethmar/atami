package test

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/stretchr/testify/assert"
)

// Create test the creator repo
func Create(t *testing.T, creator *message.Creator, newMessage message.CreateMessage) {
	message, err := creator.Create(newMessage)

	assert.Nil(t, err)
	assert.Equal(t, message.ID, 1)
	assert.Equal(t, message.Text, newMessage.Text)
	assert.Equal(t, message.CreatedByUserID, newMessage.CreatedByUserID)
	assert.True(t, time.Now().Add(time.Microsecond).After(message.CreatedAt))
}

// Create test the creator repo
func InvalidCreate(t *testing.T, creator *message.Creator, newMessage message.CreateMessage) {
	_, err := creator.Create(newMessage)
	assert.Error(t, err)
}
