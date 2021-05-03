package auth

import (
	"strings"

	"github.com/dwethmar/atami/pkg/auth/validate"
	userValidator "github.com/dwethmar/atami/pkg/user/validate"
)

// Validator struct definition
type Validator struct {
	usernameValidator *userValidator.UsernameValidator
	emailValidator    *userValidator.EmailValidator
	passwordValidator *validate.PasswordValidator
}

type errValidate struct {
	Errors []error
}

func (err errValidate) Valid() bool {
	return len(err.Errors) == 0
}

func (err errValidate) Error() string {
	errors := make([]string, len(err.Errors))
	for i, e := range err.Errors {
		errors[i] = e.Error()
	}
	return strings.Join(errors, ". ")
}

// ValidateNewUser validates a new user
func (v Validator) ValidateNewUser(newUser CreateUser) error {
	err := errValidate{}

	if e := v.usernameValidator.Validate(newUser.Username); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.emailValidator.Validate(newUser.Email); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.passwordValidator.Validate(newUser.Password); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

// NewValidator creates a new validator
func NewValidator(
	usernameValidator *userValidator.UsernameValidator,
	emailValidator *userValidator.EmailValidator,
	passwordValidator *validate.PasswordValidator,
) *Validator {
	return &Validator{
		usernameValidator,
		emailValidator,
		passwordValidator,
	}
}

// NewDefaultValidator creates a new validator
func NewDefaultValidator() *Validator {
	return NewValidator(
		userValidator.NewUsernameValidator(),
		userValidator.NewEmailValidator(),
		validate.NewPasswordValidator(),
	)
}
