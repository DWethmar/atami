package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/validate"
)

var emailValidator = validate.NewEmailValidator()
var usernameValidator = validate.NewUsernameValidator()
var passwordValidator = validate.NewPasswordValidator()

var validator = auth.NewValidator(
	usernameValidator,
	emailValidator,
	passwordValidator,
)

func TestCreate(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	auth.TestRegister(t, register, auth.RegisterUser{
		Username:      "username",
		Email:         "test@test.nl",
		PlainPassword: "!Test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := auth.RegisterUser{
		Username:      "username",
		Email:         "test@test.nl",
		PlainPassword: "!Test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	auth.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := auth.RegisterUser{
		Username:      "username",
		Email:         "test@test.nl",
		PlainPassword: "!Test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	auth.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), validator, s)
	auth.TestEmptyPassword(t, register)
}
