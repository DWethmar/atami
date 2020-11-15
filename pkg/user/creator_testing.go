package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, creator *Creator, newUser CreateUser) {
	user, err := creator.Create(newUser)

	if assert.NoError(t, err) {
		assert.NotZero(t, user.ID)
		assert.NotEmpty(t, user.UID)
		assert.Equal(t, user.Email, newUser.Email)
		time.Sleep(1)
		assert.True(t, time.Now().After(user.CreatedAt))
		assert.True(t, time.Now().After(user.UpdatedAt))
	}
}

// TestDuplicateUsername check if the correct error is returned
func TestDuplicateUsername(t *testing.T, creator *Creator, newUser CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Email = "new_" + newUser.Email

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, ErrUsernameAlreadyTaken, errTwo)
}

// TestDuplicateEmail check if the correct error is returned
func TestDuplicateEmail(t *testing.T, creator *Creator, newUser CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Username = "new_" + newUser.Username

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, ErrEmailAlreadyTaken, errTwo)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T, creator *Creator) {
	_, err := creator.Create(CreateUser{
		Username:       "wow",
		Email:          "test@test.nl",
		HashedPassword: "",
	})
	assert.EqualError(t, err, ErrPwdNotSet.Error())
}
