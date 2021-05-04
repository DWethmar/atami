package user

import (
	"strings"

	"github.com/dwethmar/atami/pkg/domain/user/validate"
)

// Validator struct definition
type Validator struct {
	usernameValidator  *validate.UsernameValidator
	emailValidator     *validate.EmailValidator
	biographyValidator *validate.BiographyValidator
	passwordValidator  *validate.PasswordValidator
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

// ValidateCreateUser validates a new user
func (v Validator) ValidateCreateUser(createUser CreateUser) error {
	err := errValidate{}

	if e := v.validateUsername(createUser); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.validateEmail(createUser); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if e := v.passwordValidator.Validate(createUser.Password); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if err.Valid() {
		return nil
	}

	return err
}

// ValidateUpdateUser validates a new user
func (v Validator) ValidateUpdateUser(updateUser UpdateUser) error {
	err := errValidate{}

	if e := v.biographyValidator.Validate(updateUser.Biography); e != nil {
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

// CreateValidator creates a new validator
func CreateValidator(
	usernameValidator *validate.UsernameValidator,
	emailValidator *validate.EmailValidator,
	biographyValidator *validate.BiographyValidator,
	passwordValidator *validate.PasswordValidator,
) *Validator {
	return &Validator{
		usernameValidator,
		emailValidator,
		biographyValidator,
		passwordValidator,
	}
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return CreateValidator(
		validate.NewUsernameValidator(),
		validate.NewEmailValidator(),
		validate.NewBiographyValidator(),
		validate.NewPasswordValidator(),
	)
}
