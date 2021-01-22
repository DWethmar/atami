package util

import (
	"errors"

	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user"
)

// ToMsgUser from the memstore to message user
func ToMsgUser(i interface{}) (*message.User, error) {
	if usr, ok := i.(user.User); ok {
		return &message.User{
			ID:       usr.ID,
			UID:      usr.UID,
			Username: usr.Username,
		}, nil
	}
	return nil, errors.New("provided value is not an user")
}
