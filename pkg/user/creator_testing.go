package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, creator *Creator, newUser NewUser) {
	user, err := creator.Create(newUser)

	if assert.NoError(t, err) {
		assert.Equal(t, user.ID, ID(1))
		assert.NotEmpty(t, user.UID)
		assert.Equal(t, user.Email, newUser.Email)
		time.Sleep(1)
		assert.True(t, time.Now().After(user.CreatedAt))
		assert.True(t, time.Now().After(user.UpdatedAt))
	}
}

// TestDuplicateEmail check if the correct error is returned
func TestDuplicateEmail(t *testing.T, creator *Creator, newUser NewUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, ErrEmailAlreadyTaken, errTwo)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T, creator *Creator) {
	_, err := creator.Create(*&NewUser{
		Username: "wow",
		Email:    "test@test.nl",
	})
	assert.Equal(t, ErrPwdNotSet, err)
}
