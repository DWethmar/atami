package auth

import (
	"testing"

	"github.com/dwethmar/atami/pkg/auth/validate"
	userValidate "github.com/dwethmar/atami/pkg/user/validate"
	"github.com/stretchr/testify/assert"
)

var validUser = CreateUser{
	Username: "username",
	Email:    "test@test.nl",
	Password: "Abc123QWERRTY@#1",
}

var emailValidator = userValidate.NewEmailValidator()
var usernameValidator = userValidate.NewUsernameValidator()
var passwordValidator = validate.NewPasswordValidator()

var validator = NewValidator(
	usernameValidator,
	emailValidator,
	passwordValidator,
)

// TestValidUser tests the validate function.
func TestValidUser(t *testing.T) {
	assert.NoError(t, validator.ValidateNewUser(validUser))
}

func TestInvalidUsername(t *testing.T) {
	invalidUsername := validUser
	invalidUsername.Username = "!@#$%^&*(Iasd"
	assert.EqualError(t, validator.ValidateNewUser(invalidUsername), userValidate.ErrUsernameContainsInvalidChars.Error())

	toLongUsername := validUser
	toLongUsername.Username = "abcdefghijklmnopqrstuvwxyz"
	assert.EqualError(t, validator.ValidateNewUser(toLongUsername), userValidate.ErrUsernameToLong.Error())

	toShortUsername := validUser
	toShortUsername.Username = "a"
	assert.EqualError(t, validator.ValidateNewUser(toShortUsername), userValidate.ErrUsernameToShort.Error())
}

func TestInvalidEmail(t *testing.T) {
	wrongEmail := validUser
	wrongEmail.Email = ""

	assert.EqualError(t, validator.ValidateNewUser(wrongEmail), userValidate.ErrEmailRequired.Error())
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
