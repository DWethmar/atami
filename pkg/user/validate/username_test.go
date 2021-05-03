package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var usernameValidator = NewUsernameValidator()

func TestUsernames(t *testing.T) {
	validUsernames := []string{
		"dennis",
		"abc",
		"ABC",
		"Abc",
		"ABCD",
		"A_B_C_D",
		"abcdefghij",
		"a__",
	}

	for _, username := range validUsernames {
		assert.NoError(t, usernameValidator.Validate(username), fmt.Sprintf("username: %s", username))
	}
}

func TestInvalidUsernames(t *testing.T) {
	invalidUsernames := []string{
		"abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij",
		"a",
		"aa",
		"****",
		"@#$%^&*",
		"___",
	}

	for _, username := range invalidUsernames {
		assert.Error(t, usernameValidator.Validate(username), fmt.Sprintf("username: %s", username))
	}
}
