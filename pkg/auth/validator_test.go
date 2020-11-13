package auth

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/dwethmar/atami/pkg/validate"
	"github.com/stretchr/testify/assert"
)

var validUser = User{
	ID:        model.UserID(1),
	UID:       "asdasdasdasd",
	Username:  "username",
	Email:     "test@test.nl",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var emailValidator = validate.NewEmailValidator()
var usernameValidator = validate.NewUsernameValidator()
var passwordValidator = validate.NewPasswordValidator()

var validator = NewValidator(
	usernameValidator,
	emailValidator,
	passwordValidator,
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
	assert.NoError(t, validator.ValidateNewUser(CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "Abcdefgh123@@",
	}))
}

func TestInvalidNewUser(t *testing.T) {
	assert.Error(t, validator.ValidateNewUser(CreateUser{
		Username: "a",
		Email:    "b",
	}))
}
