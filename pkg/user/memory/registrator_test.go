package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/validate"
)

var emailValidator = validate.NewEmailValidator()
var usernameValidator = validate.NewUsernameValidator()
var validator = user.NewValidator(usernameValidator, emailValidator)

func TestCreate(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	user.TestRegister(t, register, user.NewUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := user.NewUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	user.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.NewUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	user.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	user.TestEmptyPassword(t, register)
}
