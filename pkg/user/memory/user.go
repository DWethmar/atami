package memory

import (
	"time"

	"github.com/dwethmar/atami/pkg/user"
)

// User struct declaration
type userRecord struct {
	ID        int
	UID       string
	Username  string
	Email     string
	Salt      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func recordToUser(userRecord userRecord) *user.User {
	return &user.User{
		ID:        userRecord.ID,
		UID:       userRecord.UID,
		Username:  userRecord.Username,
		Email:     userRecord.Email,
		CreatedAt: userRecord.CreatedAt,
		UpdatedAt: userRecord.UpdatedAt,
	}
}
