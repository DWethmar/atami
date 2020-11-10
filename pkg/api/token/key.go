package token

import (
	"errors"
	"os"
)

// GetAccessSecret return the access secret from env
func GetAccessSecret() ([]byte, error) {
	t := os.Getenv("ACCESS_SECRET")
	if t == "" {
		return nil, errors.New("access token is not set")
	}
	return []byte(t), nil
}
