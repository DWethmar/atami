package user

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/dwethmar/atami/pkg/validate"
)

var (
	// UsernameMinimunLength the minimun length of the username
	UsernameMinimunLength = 3
	// UsernameMaximumLength the minimun length of the username
	UsernameMaximumLength = 15
	// ErrUsernameContainsInvalidChars error used when there are invalid chars in the username.
	ErrUsernameContainsInvalidChars = errors.New("username can only contain alphanumeric characters (letters A-Z, numbers 0-9) with the exception of underscores")
	// ErrUsernameRequired error used when there is no username
	ErrUsernameRequired = errors.New("username is required")
	// ErrUsernameToShort error used when the username is to short
	ErrUsernameToShort = fmt.Errorf("username is shorter then %d characters", UsernameMinimunLength)
	// ErrUsernameToLong error used when the username is to long
	ErrUsernameToLong = fmt.Errorf("username is longer then %d characters", UsernameMaximumLength)

	// ErrEmailRequired error used when there is no email
	ErrEmailRequired = errors.New("email is required")
	// ErrEmailInvalid error is the email is not valid
	ErrEmailInvalid = errors.New("email is invalid")
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
	var err error
	len := len(user.GetUsername())
	switch {
	case len == 0:
		err = ErrUsernameRequired
	case len < UsernameMinimunLength:
		err = ErrUsernameToShort
	case len > UsernameMaximumLength:
		err = ErrUsernameToLong
	}

	if err != nil {
		return err
	}

	if r, err := regexp.Compile("^[A-Za-z0-9][A-Za-z0-9_]{1,15}$"); err == nil {
		if !r.MatchString(user.GetUsername()) {
			return ErrUsernameContainsInvalidChars
		}
	} else {
		return err
	}

	return err
}

func (v Validator) validateEmail(user hasEmail) error {
	if user.GetEmail() == "" {
		return ErrEmailRequired
	}

	if err := v.emailValidator.Validate(user.GetEmail()); err != nil {
		return ErrEmailInvalid
	}
	return nil
}

// NewValidator creates a new validator
func NewValidator(emailValidator *validate.EmailValidator) *Validator {
	return &Validator{
		emailValidator: emailValidator,
	}
}
