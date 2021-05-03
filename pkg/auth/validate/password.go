package validate

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	// PasswordMinimunLength default minimun length of the password
	PasswordMinimunLength = 8
	// PasswordMaximumLength default minimun length of the username
	PasswordMaximumLength = 50
	// ErrPasswordRequired error used when there is no username
	ErrPasswordRequired = errors.New("password is required")
	// ErrPasswordToShort error used when the username is to short
	ErrPasswordToShort = fmt.Errorf("password is shorter then %d characters", PasswordMinimunLength)
	// ErrPasswordToLong error used when the username is to long
	ErrPasswordToLong = fmt.Errorf("password is longer then %d characters", PasswordMaximumLength)
	// ErrPasswordComplexity error used when the password is too simple
	ErrPasswordComplexity = errors.New("password is too simple")
)

// PasswordValidator struct
type PasswordValidator struct {
	minimumLength int
	maximumLength int
}

func passwordComplexity(s string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}

// Validate validates a email
func (v PasswordValidator) Validate(password string) error {
	var err error
	len := len(password)

	fmt.Printf("PASSWORD: %v", password)
	fmt.Printf("PASSWORD LEN: %v", len)

	switch {
	case password == "":
		err = ErrPasswordRequired
	case len < v.minimumLength:
		err = ErrPasswordToShort
	case len > v.maximumLength:
		err = ErrPasswordToLong
	}

	if err != nil {
		return err
	}

	if !passwordComplexity(password) {
		return ErrPasswordComplexity
	}

	return err
}

// NewPasswordValidator creates new NewEmailValidator
func NewPasswordValidator() *PasswordValidator {
	return &PasswordValidator{PasswordMinimunLength, PasswordMaximumLength}
}
