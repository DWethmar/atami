package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestUpdater test the updater repo
func TestUpdater(t *testing.T, updater *Updater, newUser UpdateRequest) {
	now := time.Now().UTC()
	time.Sleep(1)
	user, err := updater.Update(1, newUser)

	if assert.NoError(t, err) {
		if assert.NotNil(t, user) {
			assert.Equal(t, user.Biography, newUser.Biography)

			time.Sleep(1)
			assert.True(t, now.Before(user.UpdatedAt))
		}
	}
}
