package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiography(t *testing.T) {
	valid := []string{
		"lorum ipsum",
		"",
		"a",
	}

	for _, v := range valid {
		assert.NoError(t, ValidateBiography(v), fmt.Sprintf("Biography: %s", v))
	}
}

func TestInvalidBiography(t *testing.T) {
	valid := []string{
		"lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum lorum ipsum",
	}

	for _, v := range valid {
		assert.Error(t, ValidateBiography(v), fmt.Sprintf("Biography: %s", v))
	}
}

func TestEmails(t *testing.T) {
	validEmails := []string{
		"hello@me.com",
		"helloHelloHello@meme.com",
	}

	for _, email := range validEmails {
		assert.NoError(t, ValidateEmail(email), fmt.Sprintf("email: %s", email))
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
		assert.Error(t, ValidateEmail(email), fmt.Sprintf("email: %s", email))
	}
}

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
		assert.NoError(t, ValidatePassword(password), fmt.Sprintf("password: %s", password))
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
		assert.Error(t, ValidatePassword(password), fmt.Sprintf("password: %s", password))
	}
}


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
		assert.NoError(t, ValidateUsername(username), fmt.Sprintf("username: %s", username))
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
		assert.Error(t, ValidateUsername(username), fmt.Sprintf("username: %s", username))
	}
}
