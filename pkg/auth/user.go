package auth

import (
	"time"

	"github.com/dwethmar/atami/pkg/model"
)

type hasUsername interface {
	GetUsername() string
}

type hasEmail interface {
	GetEmail() string
}

// User struct declaration
type User struct {
	ID        model.UserID
	UID       model.UserUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetUsername return the username
func (u User) GetUsername() string {
	return u.Username
}

// GetEmail return the email
func (u User) GetEmail() string {
	return u.Email
}

// CreateUser struct declaration
type CreateUser struct {
	Username string
	Email    string
	Password string
}

// HashedCreateUser struct declaration internal
type HashedCreateUser struct {
	Username       string
	Email          string
	HashedPassword string
}

// GetUsername return the username
func (u CreateUser) GetUsername() string {
	return u.Username
}

// GetEmail return the email
func (u CreateUser) GetEmail() string {
	return u.Email
}

// Credentials is information used to authenticate an user
type Credentials struct {
	Email    string
	Password string
}

func toUser(user *User) *model.User {
	return &model.User{
		ID:        user.ID,
		UID:       user.UID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
