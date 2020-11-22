package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessToken(t *testing.T) {
	expiresAt := time.Now().Add(time.Hour * 10).Unix()
	accessToken, err := CreateAccessToken("abc123", "abcdefgh", expiresAt)
	assert.NoError(t, err)

	if token, err := VerifyAccessToken(accessToken); err != nil || !token.Valid {
		assert.Fail(t, fmt.Sprintf("excpected token te be valid: %v %s\n", 2, accessToken))
	}
}

func TestInvalidAccessToken(t *testing.T) {
	accessToken, err := CreateAccessToken("abc123", "abcdefgh", 1605036741)
	assert.NoError(t, err)

	if _, err := VerifyAccessToken(accessToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", accessToken))
	} else {
		assert.EqualError(t, err, ErrExpiredToken.Error())
	}
}

func TestExpiredAccessToken(t *testing.T) {
	refreshToken, err := CreateRefreshToken("abc123", "abcdefgh", 667224000)
	assert.NoError(t, err)

	if _, err := VerifyAccessToken(refreshToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", refreshToken))
	} else {
		assert.EqualError(t, err, ErrExpiredToken.Error())
	}
}

func TestRefreshToken(t *testing.T) {
	expiresAt := time.Now().Add(time.Hour * 10).Unix()
	refreshToken, err := CreateRefreshToken("abc123", "abcdefgh", expiresAt)
	assert.NoError(t, err)

	if token, err := VerifyRefreshToken(refreshToken); err != nil || !token.Valid {
		assert.Fail(t, fmt.Sprintf("excpected token te be valid: %v %s\n", 2, refreshToken))
	}
}

func TestInvalidRefreshToken(t *testing.T) {
	refreshToken, err := CreateRefreshToken("abc123", "abcdefgh", 1605036741)
	assert.NoError(t, err)

	if _, err := VerifyRefreshToken(refreshToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", refreshToken))
	} else {
		assert.EqualError(t, err, ErrExpiredToken.Error())
	}
}

func TestExpiredRefreshToken(t *testing.T) {
	refreshToken, err := CreateRefreshToken("abc123", "abcdefgh", 667224000)
	assert.NoError(t, err)

	if _, err := VerifyRefreshToken(refreshToken); err == nil {
		assert.Fail(t, fmt.Sprintf("excpected error %s \n", refreshToken))
	} else {
		assert.EqualError(t, err, ErrExpiredToken.Error())
	}
}
