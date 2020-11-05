package user

import (
	"time"

	"github.com/dwethmar/atami/pkg/auth"
)

// Responds struct declaration
type Responds struct {
	UID       string
	Username  string
	CreatedAt time.Time
}

func toResponds(users []*auth.User) []*Responds {
	r := make([]*Responds, len(users))
	for i, user := range users {
		r[i] = toRespond(user)
	}
	return r
}

func toRespond(user *auth.User) *Responds {
	return &Responds{
		UID:       user.UID.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
