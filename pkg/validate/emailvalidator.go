package validate

import "errors"

var (
	// ErrEmailRequired error used when there is no email
	ErrEmailRequired = errors.New("email is required")
	// ErrEmailInvalid error is the email is not valid
	ErrEmailInvalid = errors.New("email is invalid")
)

// EmailValidator struct
type EmailValidator struct {
	minimumLength int
	maximumLength int
}

// Validate validates a email
func (v EmailValidator) Validate(email string) error {
	if email == "" {
		return ErrEmailRequired
	}
	if len(email) < v.minimumLength || len(email) > v.maximumLength {
		return ErrEmailInvalid
	}
	return nil
}

// NewEmailValidator creates new NewEmailValidator
func NewEmailValidator() *EmailValidator {
	return &EmailValidator{5, 254}
}
