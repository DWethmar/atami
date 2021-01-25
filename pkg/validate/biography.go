package validate

import (
	"fmt"
)

var (
	// BiographyMaximumLength default minimun length of the Biography
	BiographyMaximumLength = 250
	// ErrBiographyToLong error used when the Biography is to long
	ErrBiographyToLong = fmt.Errorf("biography is longer then %d characters", BiographyMaximumLength)
)

// BiographyValidator struct
type BiographyValidator struct {
	maximumLength int
}

// Validate validates a email
func (v BiographyValidator) Validate(biography string) error {
	if len(biography) > v.maximumLength {
		return ErrUsernameToLong
	}
	return nil
}

// NewBiographyValidator creates new NewEmailValidator
func NewBiographyValidator() *BiographyValidator {
	return &BiographyValidator{BiographyMaximumLength}
}
