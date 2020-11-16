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
	SessionID int64 `json:"sid,omitempty"`
	jwt.StandardClaims
}

// Details contains details about a generated token
type Details struct {
	AccessToken        string
	AccessTokenExpires int64
}

func getAccessSecret() ([]byte, error) {
	t := config.Load().AccessSecret
	if t == "" {
		return nil, errors.New("access secret is not set in env: ACCESS_SECRET")
	}
	return []byte(t), nil
}

// CreateToken creates a new authentication token
func CreateToken(UID string, username string, expiresOn int64) (*Details, error) {
	td := &Details{}

	td.AccessTokenExpires = expiresOn

	var err error
	claims := CustomClaims{
		SessionID: time.Now().Unix(),
	}
	claims.StandardClaims = jwt.StandardClaims{
		Subject:   UID,
		ExpiresAt: td.AccessTokenExpires,
		IssuedAt:  time.Now().Unix(),
	}

	accessSecret, err := getAccessSecret()
	if err != nil {
		return nil, err
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// VerifyToken verifies the token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	var claims CustomClaims
	return jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if time.Now().Unix() > claims.ExpiresAt {
			return nil, ErrExpiredToken
		}

		if claims.SessionID == 0 {
			return nil, errors.New("empty session ID")
		}

		return getAccessSecret()
	})
}
