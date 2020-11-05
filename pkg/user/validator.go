package user

import (
	"strings"

	"github.com/dwethmar/atami/pkg/validate"
)

// Validator struct definition
type Validator struct {
	usernameValidator *validate.UsernameValidator
	emailValidator    *validate.EmailValidator
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

// ValidateUser validates a user
func (v Validator) ValidateUser(user User) error {
	err := errValidate{}

	if e := v.validateUsername(user); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.validateEmail(user); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

// ValidateNewUser validates a new user
func (v Validator) ValidateNewUser(newUser NewUser) error {
	err := errValidate{}

	if e := v.validateUsername(newUser); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.validateEmail(newUser); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

func (v Validator) validateUsername(user hasUsername) error {
	if err := v.usernameValidator.Validate(user.GetUsername()); err != nil {
		return err
	}
	return nil
}

func (v Validator) validateEmail(user hasEmail) error {
	if err := v.emailValidator.Validate(user.GetEmail()); err != nil {
		return err
	}
	return nil
}

// NewValidator creates a new validator
func NewValidator(
	usernameValidator *validate.UsernameValidator,
	emailValidator *validate.EmailValidator,
) *Validator {
	return &Validator{
		usernameValidator,
		emailValidator,
	}
}
