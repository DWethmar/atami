package user

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/validate"
	"github.com/stretchr/testify/assert"
)

var validUser = User{
	ID:        ID(1),
	UID:       "asdasdasdasd",
	Username:  "username",
	Email:     "test@test.nl",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

// TestValidUser tests the validate function.
func TestValidUser(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	assert.NoError(t, validator.ValidateUser(validUser))
}

func TestInvalidUsername(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	wrongEmail := validUser
	wrongEmail.Username = "!@#$%^&*(Iasd"
	assert.EqualError(t, validator.ValidateUser(wrongEmail), ErrUsernameContainsInvalidChars.Error())

	toLongUsername := validUser
	toLongUsername.Username = "abcdefghijklmnopqrstuvwxyz"
	assert.EqualError(t, validator.ValidateUser(toLongUsername), ErrUsernameToLong.Error())

	toShortUsername := validUser
	toShortUsername.Username = "a"
	assert.EqualError(t, validator.ValidateUser(toShortUsername), ErrUsernameToShort.Error())
}

func TestInvalidEmail(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	wrongEmail := validUser
	wrongEmail.Email = ""

	assert.EqualError(t, validator.ValidateUser(wrongEmail), ErrEmailRequired.Error())
}

func TestValidNewUser(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	assert.NoError(t, validator.ValidateNewUser(NewUser{
		Username: "username",
		Email:    "test@test.nl",
	}))
}

func TestInvalidNewUser(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	assert.Error(t, validator.ValidateNewUser(NewUser{
		Username: "a",
		Email:    "b",
	}))
}
