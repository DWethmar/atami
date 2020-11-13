package token

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "test123")

	details, err := CreateToken(auth.UID("abc123"), "username", time.Now().Add(time.Hour*10).Unix())
	assert.NoError(t, err)

	if token, err := VerifyToken(details.AccessToken); err != nil || !token.Valid {
		assert.Fail(t, fmt.Sprintf("excpected token te be valid: %v %s\n", 2, details.AccessToken))
	}
}

func TestInvalidToken(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "test123")

	details, err := CreateToken(auth.UID("abc123"), "username", 1605036741)
	assert.NoError(t, err)

	if _, err := VerifyToken(details.AccessToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", details.AccessToken))
	} else {
		assert.EqualError(t, err, "Token is expired")
	}
}
