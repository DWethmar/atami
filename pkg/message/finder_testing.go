package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindOne tests the ReadOne function.
func TestFindOne(t *testing.T, finder *Finder, ID int, message Message) {
	m, err := finder.FindByID(ID)
	assert.NoError(t, err)

	assert.NotEmpty(t, message.ID)
	assert.NotEmpty(t, message.UID)
	assert.NotEmpty(t, message.Text)
	assert.NotEmpty(t, message.CreatedByUserID)
	assert.False(t, message.CreatedAt.IsZero())

	if assert.NotNil(t, m) {
		assert.NotEmpty(t, m.ID)
		assert.NotEmpty(t, m.UID)
		assert.NotEmpty(t, m.Text)
		assert.NotEmpty(t, m.CreatedByUserID)
		assert.False(t, m.CreatedAt.IsZero())

		assert.Equal(t, message.ID, m.ID)
		assert.Equal(t, message.Text, m.Text)
		assert.Equal(t, message.CreatedByUserID, m.CreatedByUserID)
	}
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(0)
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestFind tests the Find function.
func TestFind(t *testing.T, finder *Finder, length int, messages []Message) {
	list, err := finder.Find(0, length)

	assert.NoError(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, message := range list {
			assert.NotEmpty(t, messages[i].ID)
			assert.NotEmpty(t, messages[i].UID)
			assert.NotEmpty(t, messages[i].Text)
			assert.NotEmpty(t, messages[i].CreatedByUserID)
			assert.False(t, messages[i].CreatedAt.IsZero())

			assert.Equal(t, messages[i].ID, message.ID)
			assert.Equal(t, messages[i].UID, message.UID)
			assert.Equal(t, messages[i].Text, message.Text)

			fmt.Printf("HMMMMM")
			fmt.Print(messages[i])

			assert.NotNil(t, messages[i].User)
		}
	}
}
