package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	details, err := CreateToken(model.UserUID("abc123"), "username", time.Now().Add(time.Hour*10).Unix())
	assert.NoError(t, err)

	if token, err := VerifyToken(details.AccessToken); err != nil || !token.Valid {
		assert.Fail(t, fmt.Sprintf("excpected token te be valid: %v %s\n", 2, details.AccessToken))
	}
}

func TestInvalidToken(t *testing.T) {
	details, err := CreateToken(model.UserUID("abc123"), "username", 1605036741)
	assert.NoError(t, err)

	if _, err := VerifyToken(details.AccessToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", details.AccessToken))
	} else {
		assert.EqualError(t, err, "Token is expired")
	}
}
