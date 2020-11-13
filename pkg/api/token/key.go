package token

import (
	"errors"

	"github.com/dwethmar/atami/pkg/config"
)

// GetAccessSecret return the access secret from env
func GetAccessSecret() ([]byte, error) {
	t := config.Load().AccessSecret
	if t == "" {
		return nil, errors.New("access token is not set")
	}
	return []byte(t), nil
}
