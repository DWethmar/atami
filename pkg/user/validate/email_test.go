package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var emailValidator = NewEmailValidator()

func TestEmails(t *testing.T) {
	validEmails := []string{
		"hello@me.com",
		"helloHelloHello@meme.com",
	}

	for _, email := range validEmails {
		assert.NoError(t, emailValidator.Validate(email), fmt.Sprintf("email: %s", email))
	}
}

func TestInvalidEmails(t *testing.T) {
	invalidEmails := []string{
		"com",
		"hey",
		"a",
		"",
	}

	for _, email := range invalidEmails {
		assert.Error(t, emailValidator.Validate(email), fmt.Sprintf("email: %s", email))
	}
}
