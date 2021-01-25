package node

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, creator *Creator, newNode CreateRequest) {
	node, err := creator.Create(newNode)

	assert.Nil(t, err)
	assert.Equal(t, node.ID, 1)
	assert.Equal(t, node.Text, newNode.Text)
	assert.Equal(t, node.CreatedByUserID, newNode.CreatedByUserID)
	assert.True(t, time.Now().Add(time.Microsecond).After(node.CreatedAt))
}
