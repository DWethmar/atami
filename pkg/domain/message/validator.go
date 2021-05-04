package message

import (
	"errors"
	"strings"

	"github.com/dwethmar/atami/pkg/domain/message/validate"
)

// Validator struct definition
type Validator struct {
	messageTextValidator *validate.MessageTextValidator
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

// ValidateCreate validates a new message
func (v Validator) ValidateCreate(msg CreateMessage) error {
	err := errValidate{}

	if e := v.messageTextValidator.Validate(msg.Text); e != nil {
		err.Errors = append(err.Errors, e)
	}

	if msg.CreatedByUserID == 0 {
		err.Errors = append(err.Errors, errors.New("user not set"))
	}

	if err.Valid() {
		return nil
	}

	return err
}

// NewValidator creates a new validator
func NewValidator(
	messageTextValidator *validate.MessageTextValidator,
) *Validator {
	return &Validator{
		messageTextValidator: messageTextValidator,
	}
}

// NewDefaultValidator creates a new validator
func NewDefaultValidator() *Validator {
	return NewValidator(
		validate.NewMessageTextValidator(),
	)
}
