package node

import (
	"errors"
	"strings"

	"github.com/dwethmar/atami/pkg/validate"
)

// Validator struct definition
type Validator struct {
	nodeTextValidator *validate.MessageTextValidator
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

// ValidateCreateNode validates a new node
func (v Validator) ValidateCreateNode(msg CreateRequest) error {
	err := errValidate{}

	if e := v.nodeTextValidator.Validate(msg.Text); e != nil {
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
	nodeTextValidator *validate.MessageTextValidator,
) *Validator {
	return &Validator{
		nodeTextValidator: nodeTextValidator,
	}
}

// NewDefaultValidator creates a new validator
func NewDefaultValidator() *Validator {
	return NewValidator(
		validate.NewMessageTextValidator(),
	)
}
