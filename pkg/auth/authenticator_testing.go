package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAuthenticate tests the Authenticate function.
func TestAuthenticate(t *testing.T, authenticator *Authenticator, credentials Credentials) {
	ok, err := authenticator.Authenticate(credentials)
	assert.NoError(t, err)
	assert.True(t, ok)
}
