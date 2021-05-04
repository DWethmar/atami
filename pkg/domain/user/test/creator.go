package test

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/validate"
	"github.com/stretchr/testify/assert"
)

// Creator test the creator repo
func Creator(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
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

// DuplicateUsername check if the correct error is returned
func DuplicateUsername(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Email = "new_" + newUser.Email

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, user.ErrUsernameAlreadyTaken, errTwo)
}

// DuplicateEmail check if the correct error is returned
func DuplicateEmail(t *testing.T, creator *user.Creator, newUser user.CreateUser) {
	_, errOne := creator.Create(newUser)
	assert.NoError(t, errOne)

	newUser.Username = "new_" + newUser.Username

	_, errTwo := creator.Create(newUser)
	assert.Equal(t, user.ErrEmailAlreadyTaken, errTwo)
}

// EmptyPassword test if the correct error is returned
func EmptyPassword(t *testing.T, creator *user.Creator) {
	_, err := creator.Create(user.CreateUser{
		Username: "wow",
		Email:    "test@test.nl",
		Password: "",
	})
	assert.EqualError(t, err, validate.ErrPasswordRequired.Error())
}
