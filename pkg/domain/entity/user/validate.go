package user

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

var (
	// BiographyMaximumLength default minimun length of the Biography
	BiographyMaximumLength = 250
	// ErrBiographyToLong error used when the Biography is to long
	ErrBiographyToLong = fmt.Errorf("biography is longer then %d characters", BiographyMaximumLength)

	// EmailMinimumLength default minimun length of email addresses
	EmailMinimumLength = 4
	// EmailMinimumLength default minimun length of email addresses
	EmailMaximumLength = 254
	// ErrEmailRequired error used when there is no email
	ErrEmailRequired = errors.New("email is required")
	// ErrEmailInvalid error is the email is not valid
	ErrEmailInvalid = errors.New("email is invalid")

	// PasswordMinimunLength default minimun length of the password
	PasswordMinimunLength = 8
	// PasswordMaximumLength default maximum length of the password
	PasswordMaximumLength = 50
	// ErrPasswordRequired error used when there is no username
	ErrPasswordRequired = errors.New("password is required")
	// ErrPasswordToShort error used when the username is to short
	ErrPasswordToShort = fmt.Errorf("password is shorter then %d characters", PasswordMinimunLength)
	// ErrPasswordToLong error used when the username is to long
	ErrPasswordToLong = fmt.Errorf("password is longer then %d characters", PasswordMaximumLength)
	// ErrPasswordComplexity error used when the password is too simple
	ErrPasswordComplexity = errors.New("password is too simple")

	// UsernameMinimunLength default minimun length of the username
	UsernameMinimunLength = 3
	// UsernameMaximumLength default minimun length of the username
	UsernameMaximumLength = 15
	// ErrUsernameContainsInvalidChars error used when there are invalid chars in the username.
	ErrUsernameContainsInvalidChars = errors.New("username can only contain alphanumeric characters (letters A-Z, numbers 0-9) with the exception of underscores")
	// ErrUsernameRequired error used when there is no username
	ErrUsernameRequired = errors.New("username is required")
	// ErrUsernameToShort error used when the username is to short
	ErrUsernameToShort = fmt.Errorf("username is shorter then %d characters", UsernameMinimunLength)
	// ErrUsernameToLong error used when the username is to long
	ErrUsernameToLong = fmt.Errorf("username is longer then %d characters", UsernameMaximumLength)
)

// Validate validates a email
func ValidateBiography(biography string) error {
	if len(biography) > BiographyMaximumLength {
		return ErrBiographyToLong
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}
	if len(email) < EmailMinimumLength || len(email) > EmailMaximumLength {
		return ErrEmailInvalid
	}
	return nil
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
func ValidatePassword(password string) error {
	var err error
	len := len(password)

	switch {
	case password == "":
		err = ErrPasswordRequired
	case len < PasswordMinimunLength:
		err = ErrPasswordToShort
	case len > PasswordMaximumLength:
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

func ValidateUsername(username string) error {
	var err error
	len := len(username)
	switch {
	case username == "":
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
		if !r.MatchString(username) {
			return ErrUsernameContainsInvalidChars
		}
	} else {
		return err
	}

	return err
}