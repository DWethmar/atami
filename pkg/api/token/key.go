package token

import (
	"errors"
	"os"
)

// GetAccessSecret return the access secret from env
func GetAccessSecret() (string, error) {
	t := os.Getenv("ACCESS_SECRET")
	if t == "" {
		return "", errors.New("access token is not set")
	}
	return t, nil
}
