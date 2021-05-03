package validate

import (
	"errors"
	"fmt"
)

var (
	// MsgTxtMinimunLength default minimun length of the username
	MsgTxtMinimunLength = 3
	// MsgTxtMaximumLength default minimun length of the username
	MsgTxtMaximumLength = 300
	// ErrMsgTxtRequired error used when there is no text
	ErrMsgTxtRequired = errors.New("text is required")
	// ErrMsgTxtInvalid error is the text is not valid
	ErrMsgTxtInvalid = errors.New("text is invalid")
	// ErrMsgTxtExceedMaxLength error when text exceed  max length
	ErrMsgTxtExceedMaxLength = fmt.Errorf("message exceeds max length of %v", MsgTxtMaximumLength)
	// ErrMsgTxtFailsMinLength error when text fails  max length
	ErrMsgTxtFailsMinLength = fmt.Errorf("message fails min length of %v", MsgTxtMinimunLength)
)

// MessageTextValidator struct
type MessageTextValidator struct {
	minimumLength int
	maximumLength int
}

// Validate validates a email
func (v MessageTextValidator) Validate(txt string) error {
	if txt == "" {
		return ErrMsgTxtRequired
	}
	if len(txt) < v.minimumLength || len(txt) > v.maximumLength {
		return ErrMsgTxtInvalid
	}
	return nil
}

// NewMessageTextValidator creates new MessageTextValidator
func NewMessageTextValidator() *MessageTextValidator {
	return &MessageTextValidator{
		MsgTxtMinimunLength,
		MsgTxtMaximumLength,
	}
}
