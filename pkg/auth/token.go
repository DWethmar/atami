package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dwethmar/atami/pkg/config"
)

// https://www.sohamkamani.com/golang/2019-01-01-jwt-authentication/

// ErrExpiredToken defines a error
var ErrExpiredToken = errors.New("token is expired")

// CustomClaims adds soem fields to the standard claims
type CustomClaims struct {
	SessionID    string `json:"sid,omitempty"`
	AllowRefresh string `json:"ref,omitempty"`
	jwt.StandardClaims
}

func getAccessSecret() ([]byte, error) {
	t := config.Load().AccessSecret
	if t == "" {
		return nil, errors.New("access secret is not set in env: ACCESS_SECRET")
	}
	return []byte(t), nil
}

// CreateAccessToken creates a new authentication token
func CreateAccessToken(UID string, session string, expiresOn int64) (string, error) {
	var err error

	accessSecret, err := getAccessSecret()
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		SessionID:    session,
		AllowRefresh: "0",
	}

	claims.StandardClaims = jwt.StandardClaims{
		Subject:   UID,
		ExpiresAt: expiresOn,
		IssuedAt:  time.Now().Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(accessSecret))
}

// CreateRefreshToken creates a new authentication token
func CreateRefreshToken(UID string, session string, expiresOn int64) (string, error) {
	var err error

	accessSecret, err := getAccessSecret()
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		SessionID:    session,
		AllowRefresh: "1",
	}

	claims.StandardClaims = jwt.StandardClaims{
		Subject:   UID,
		ExpiresAt: expiresOn,
		IssuedAt:  time.Now().Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(accessSecret))
}

// VerifyAccessToken verifies the access token
func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	var claims CustomClaims
	return jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if time.Now().Unix() > claims.ExpiresAt {
			return nil, ErrExpiredToken
		}

		if claims.SessionID == "" {
			return nil, errors.New("empty session ID")
		}

		if claims.AllowRefresh == "1" {
			return nil, errors.New("token is no access token")
		}

		return getAccessSecret()
	})
}

// VerifyRefreshToken verifies the refresh token
func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	var claims CustomClaims
	return jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if time.Now().Unix() > claims.ExpiresAt {
			return nil, ErrExpiredToken
		}

		if claims.SessionID == "" {
			return nil, errors.New("empty session ID")
		}

		if claims.AllowRefresh == "0" {
			return nil, errors.New("token is no refresh token")
		}

		return getAccessSecret()
	})
}
