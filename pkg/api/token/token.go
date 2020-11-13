package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dwethmar/atami/pkg/model"
)

// https://www.sohamkamani.com/golang/2019-01-01-jwt-authentication/

// Details contains details about a generated token
type Details struct {
	AccessToken        string
	AccessTokenExpires int64
}

// CreateToken creates a new authentication token
func CreateToken(UID model.UserUID, username string, expiresOn int64) (*Details, error) {
	td := &Details{}

	td.AccessTokenExpires = expiresOn

	var err error
	//Creating Access Token
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = UID
	claims["exp"] = td.AccessTokenExpires
	claims["iat"] = time.Now().Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessSecret, err := GetAccessSecret()
	if err != nil {
		return nil, err
	}

	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// VerifyToken verifies the token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return GetAccessSecret()
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
