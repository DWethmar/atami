package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUpdater test the updater repo
func TestUpdater(t *testing.T, updater *Updater, newUser UpdateRequest) {
	user, err := updater.Update(newUser)

	if assert.NoError(t, err) {
		assert.Equal(t, user.Biography, newUser.Biography)
		assert.True(t, user.CreatedAt.Before(user.UpdatedAt))
	}
}
