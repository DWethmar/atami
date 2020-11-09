package token

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessSecret(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "abc")
	secret, err := GetAccessSecret()
	assert.NoError(t, err)
	assert.NotEmpty(t, secret)
}
