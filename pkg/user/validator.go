package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dwethmar/atami/pkg/validate"
)

// Validator struct definition
type Validator struct {
	emailValidator *validate.EmailValidator
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

	if e := v.validateName(user); e != nil {
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

	if e := v.validateName(newUser); e != nil {
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

func (v Validator) validateName(user hasUsername) error {
	len := len(user.GetUsername())
	var err error
	switch {
	case len == 0:
		err = errors.New("username is empty")
	case len < 3:
		err = fmt.Errorf("username is shorter then %d characters", 3)
	case len > 12:
		err = fmt.Errorf("username is larger then %d characters", 12)
	}

	if err == nil {
		return nil
	}

	return err
}

func (v Validator) validateEmail(user hasEmail) error {
	if v.emailValidator.Validate(user.GetEmail()) {
		return nil
	}
	return errors.New("Invalid email")
}

// NewValidator creates a new validator
func NewValidator(emailValidator *validate.EmailValidator) *Validator {
	return &Validator{
		emailValidator: emailValidator,
	}
}
