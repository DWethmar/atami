package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestRegister test the creator repo
func TestRegister(t *testing.T, register *Registrator, newUser RegisterUser) {
	user, err := register.Register(newUser)

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
func TestDuplicateUsername(t *testing.T, register *Registrator, newUser RegisterUser) {
	_, errOne := register.Register(newUser)
	assert.NoError(t, errOne)

	newUser.Email = "new_" + newUser.Email

	_, errTwo := register.Register(newUser)
	assert.Equal(t, ErrUsernameAlreadyTaken, errTwo)
}

// TestDuplicateEmail check if the correct error is returned
func TestDuplicateEmail(t *testing.T, register *Registrator, newUser RegisterUser) {
	_, errOne := register.Register(newUser)
	assert.NoError(t, errOne)

	newUser.Username = "new_" + newUser.Username

	_, errTwo := register.Register(newUser)
	assert.Equal(t, ErrEmailAlreadyTaken, errTwo)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T, register *Registrator) {
	_, err := register.Register(RegisterUser{
		Username: "wow",
		Email:    "test@test.nl",
		Password: "",
	})
	assert.EqualError(t, err, "password is required")
}
