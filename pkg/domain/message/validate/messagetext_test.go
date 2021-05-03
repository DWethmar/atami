package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var msgTxtValidator = NewMessageTextValidator()

func TestMessageText(t *testing.T) {
	validTxt := []string{
		"test",
	}

	for _, email := range validTxt {
		assert.NoError(t, msgTxtValidator.Validate(email), fmt.Sprintf("txt: %s", email))
	}
}

func TestInvalidMessageText(t *testing.T) {
	invalidTxt := []string{
		"",
	}

	for _, email := range invalidTxt {
		assert.Error(t, msgTxtValidator.Validate(email), fmt.Sprintf("txt: %s", email))
	}
}
