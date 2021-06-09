package message

import (
	"errors"
	"fmt"

	"github.com/dwethmar/atami/pkg/msgtxt"
)
var (
	// MsgTxtMinimunLength default minimun length of the username
	MsgTxtMinimunLength = 3
	// MsgTxtMaximumLength default minimun length of the username
	MsgTxtMaximumLength = 300
	// ErrMsgTxtRequired error used when there is no text
	ErrMsgTxtRequired = errors.New("text is required")
	// ErrMsgTxtExceedMaxLength error when text exceed  max length
	ErrMsgTxtExceedMaxLength = fmt.Errorf("message exceeds max length of %v characters", MsgTxtMaximumLength)
	// ErrMsgTxtFailsMinLength error when text fails  max length
	ErrMsgTxtFailsMinLength = fmt.Errorf("message fails min length of %v characters", MsgTxtMinimunLength)
)

func ValidateText(txt string) error {
	r := msgtxt.Parse(txt)

	if r.NormalizedLength == 0 {
		return ErrMsgTxtRequired
	}

	if r.NormalizedLength < MsgTxtMinimunLength{
		return ErrMsgTxtFailsMinLength
	}

	if r.NormalizedLength > MsgTxtMaximumLength {
		return ErrMsgTxtExceedMaxLength
	}

	return nil
}