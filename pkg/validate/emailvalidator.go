package validate

// EmailValidator struct
type EmailValidator struct {
	minimumLength int
	maximumLength int
}

// Validate validates a email
func (v EmailValidator) Validate(email string) bool {
	return len(email) > v.minimumLength && len(email) < v.maximumLength
}

// NewEmailValidator creates new NewEmailValidator
func NewEmailValidator() *EmailValidator {
	return &EmailValidator{3, 64}
}
