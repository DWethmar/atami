package user

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/validate"
	"github.com/stretchr/testify/assert"
)

// TestValidUser tests the validate function.
func TestValidUser(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	assert.NoError(t, validator.ValidateUser(User{
		ID:        ID(1),
		UID:       "asdasdasdasd",
		Username:  "username",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}))
}

func TestInvalidUser(t *testing.T) {
	emailValidator := validate.NewEmailValidator()
	validator := NewValidator(emailValidator)

	assert.Error(t, validator.ValidateUser(User{
		ID:        ID(1),
		UID:       "a",
		Username:  "a",
		Email:     "b",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}))
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
