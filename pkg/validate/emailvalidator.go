package validate

import "errors"

// EmailValidator struct
type EmailValidator struct {
	minimumLength int
	maximumLength int
}

// Validate validates a email
func (v EmailValidator) Validate(email string) error {
	if len(email) < v.minimumLength || len(email) > v.maximumLength {
		return errors.New("Invalid email")
	}
	return nil
}

// NewEmailValidator creates new NewEmailValidator
func NewEmailValidator() *EmailValidator {
	return &EmailValidator{3, 64}
}
