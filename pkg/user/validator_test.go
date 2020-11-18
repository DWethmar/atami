package user

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/validate"
	"github.com/stretchr/testify/assert"
)

var validUser = User{
	ID:        1,
	UID:       "asdasdasdasd",
	Username:  "username",
	Email:     "test@test.nl",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var emailValidator = validate.NewEmailValidator()
var usernameValidator = validate.NewUsernameValidator()

var validator = NewValidator(
	usernameValidator,
	emailValidator,
)

// TestValidUser tests the validate function.
func TestValidUser(t *testing.T) {
	assert.NoError(t, validator.ValidateUser(validUser))
}

func TestInvalidUsername(t *testing.T) {
	invalidUsername := validUser
	invalidUsername.Username = "!@#$%^&*(Iasd"
	assert.EqualError(t, validator.ValidateUser(invalidUsername), validate.ErrUsernameContainsInvalidChars.Error())

	toLongUsername := validUser
	toLongUsername.Username = "abcdefghijklmnopqrstuvwxyz"
	assert.EqualError(t, validator.ValidateUser(toLongUsername), validate.ErrUsernameToLong.Error())

	toShortUsername := validUser
	toShortUsername.Username = "a"
	assert.EqualError(t, validator.ValidateUser(toShortUsername), validate.ErrUsernameToShort.Error())
}

func TestInvalidEmail(t *testing.T) {
	wrongEmail := validUser
	wrongEmail.Email = ""

	assert.EqualError(t, validator.ValidateUser(wrongEmail), validate.ErrEmailRequired.Error())
}

func TestValidNewUser(t *testing.T) {
	assert.NoError(t, validator.ValidateCreateUser(CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "Abcdefgh123@@",
	}))
}

func TestInvalidNewUser(t *testing.T) {
	assert.Error(t, validator.ValidateCreateUser(CreateUser{
		Username: "a",
		Email:    "b",
	}))
}
