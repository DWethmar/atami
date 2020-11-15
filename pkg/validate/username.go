package validate

import (
	"errors"
	"fmt"
	"regexp"
)

var (
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

// UsernameValidator struct
type UsernameValidator struct {
	minimumLength int
	maximumLength int
}

// Validate validates a email
func (v UsernameValidator) Validate(username string) error {
	var err error
	len := len(username)
	switch {
	case username == "":
		err = ErrUsernameRequired
	case len < v.minimumLength:
		err = ErrUsernameToShort
	case len > v.maximumLength:
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

// NewUsernameValidator creates new NewEmailValidator
func NewUsernameValidator() *UsernameValidator {
	return &UsernameValidator{UsernameMinimunLength, UsernameMaximumLength}
}
