package user

import (
	"time"
)

type hasUsername interface {
	GetUsername() string
}

type hasEmail interface {
	GetEmail() string
}

// User struct declaration
type User struct {
	ID        int
	UID       string
	Username  string
	Email     string
	Password  string
	Biography string
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

// Equal check if a user is equal
func (u User) Equal(user User) bool {
	return (u.ID == user.ID &&
		u.UID == user.UID &&
		u.Email == user.Email &&
		u.CreatedAt.Equal(user.CreatedAt) &&
		u.UpdatedAt.Equal(user.UpdatedAt))
}

// CreateUser struct declaration
type CreateUser struct {
	UID       string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetUsername return the username
func (u CreateUser) GetUsername() string {
	return u.Username
}

// GetEmail return the email
func (u CreateUser) GetEmail() string {
	return u.Email
}

// UpdateUser struct declaration
type UpdateUser struct {
	Biography string
	UpdatedAt time.Time
}
