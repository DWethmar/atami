package test

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

// TestCreator test the creator repo
func TestCreator(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
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
func TestDuplicateUsername(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Email = "new_" + newUser.Email

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, user.ErrUsernameAlreadyTaken, errTwo)
}

// TestDuplicateEmail check if the correct error is returned
func TestDuplicateEmail(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Username = "new_" + newUser.Username

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, user.ErrEmailAlreadyTaken, errTwo)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T, creator *user.Creator) {
	_, err := creator.Create(user.CreateUser{
		Username: "wow",
		Email:    "test@test.nl",
		Password: "",
	})
	assert.EqualError(t, err, user.ErrPwdNotSet.Error())
}
