package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var passwordvalidator = NewPasswordValidator()

func TestPasswords(t *testing.T) {
	validPasswords := []string{
		"Asd123@111",
		"9<w4Ge5Z9@rx~^}L",
		"8uFTV5H><r=L:N'm",
		"J4pT3^5P':7Fg#KJXL`;>",
		"@1Abcdef",
		"K5&.,MSn2?%=E-9-&R}GqS#pCw7dp,qFL[m6@bRt!Jn2Ayg*Me",
	}

	for _, password := range validPasswords {
		assert.NoError(t, passwordvalidator.Validate(password), fmt.Sprintf("password: %s", password))
	}
}

func TestInvalidPasswords(t *testing.T) {
	invalidPasswords := []string{
		"com",
		"hello",
		"a",
		"",
		"abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij",
	}

	for _, password := range invalidPasswords {
		assert.Error(t, passwordvalidator.Validate(password), fmt.Sprintf("password: %s", password))
	}
}
