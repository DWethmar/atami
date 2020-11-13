package postgres

import (
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/model"
)

var (
	tableName = "public.user"
)

// User struct declaration
type userRecord struct {
	ID        model.UserID
	UID       model.UserUID
	Username  string
	Email     string
	Salt      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func recordToUser(user userRecord) *auth.User {
	return &auth.User{
		ID:        user.ID,
		UID:       user.UID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
