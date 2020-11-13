package message

import (
	"testing"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

// TestFindOne tests the ReadOne function.
func TestFindOne(t *testing.T, finder *Finder, ID model.MessageID, message model.Message) {
	m, err := finder.FindByID(ID)
	assert.NoError(t, err)

	if assert.NotNil(t, m) {
		assert.Equal(t, message.ID, m.ID)
		assert.Equal(t, message.Text, m.Text)
		assert.Equal(t, message.CreatedBy, m.CreatedBy)
		assert.False(t, m.CreatedAt.IsZero())
	}
}

// TestNotFound tests the ReadOne function for a not found error.
func TestNotFound(t *testing.T, finder *Finder) {
	_, err := finder.FindByID(model.MessageID(0))
	assert.Equal(t, ErrCouldNotFind, err)
}

// TestFindAll tests the ReadOne function.
func TestFindAll(t *testing.T, finder *Finder, length int, messages []model.Message) {
	list, err := finder.FindAll()

	assert.NoError(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, message := range list {
			assert.Equal(t, messages[i].ID, message.ID)
			assert.Equal(t, messages[i].UID, message.UID)
			assert.Equal(t, messages[i].Text, message.Text)
		}
	}
}
