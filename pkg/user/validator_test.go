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
var biographyValidator = validate.NewBiographyValidator()

var validator = NewValidator(
	usernameValidator,
	emailValidator,
	biographyValidator,
)

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
