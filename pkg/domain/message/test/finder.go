package test

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/stretchr/testify/assert"
)

// FindByID tests the ReadOne function.
func FindByID(t *testing.T, finder *message.Finder, ID int, message message.Message) {
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

		if assert.NotNil(t, m.User) {
			assert.Equal(t, message.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// FindByUID tests the findByUID function.
func FindByUID(t *testing.T, finder *message.Finder, UID string, message message.Message) {
	m, err := finder.FindByUID(UID)
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

		if assert.NotNil(t, m.User) {
			assert.Equal(t, message.CreatedByUserID, m.User.ID)
			assert.Equal(t, m.CreatedByUserID, m.User.ID)
		}
	}
}

// NotFound tests the ReadOne function for a not found error.
func NotFound(t *testing.T, finder *message.Finder) {
	_, err := finder.FindByID(0)
	assert.Equal(t, message.ErrCouldNotFind, err)
}

// Find tests the Find function.
func Find(t *testing.T, finder *message.Finder, length int, messages []message.Message) {
	list, err := finder.Find(0, length)

	assert.NoError(t, err)
	if assert.Equal(t, length, len(list)) {
		for i, message := range list {
			assert.NotEmpty(t, messages[i].ID)
			assert.NotEmpty(t, messages[i].UID)
			assert.NotEmpty(t, messages[i].Text)
			assert.NotZero(t, messages[i].CreatedByUserID)
			assert.False(t, messages[i].CreatedAt.IsZero())

			assert.Equal(t, messages[i].ID, message.ID)
			assert.Equal(t, messages[i].UID, message.UID)
			assert.Equal(t, messages[i].Text, message.Text)

			if assert.NotNil(t, message.User) {
				// fmt.Println(message)
				// fmt.Println("UID: -> " + message.User.UID + " <- ")
				// fmt.Println(message.User.ID)
				// fmt.Println(message.User.Username)
				assert.Equal(t, message.CreatedByUserID, message.User.ID)
			}
		}
	}
}
