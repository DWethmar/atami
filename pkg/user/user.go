package user

import (
	"strconv"
	"time"
)

// ID type the id type used for users
type ID int64

func (ID ID) String() string {
	return strconv.FormatInt(int64(ID), 10)
}

// UID type the unique identifier for users.
type UID string

func (UID UID) String() string {
	return string(UID)
}

type hasUsername interface {
	GetUsername() string
}

type hasEmail interface {
	GetEmail() string
}

// User struct declaration
type User struct {
	ID        ID
	UID       UID
	Username  string
	Email     string
	Password  string
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

// NewUser struct declaration
type NewUser struct {
	Username string
	Email    string
	Password string
}

// GetUsername return the username
func (u NewUser) GetUsername() string {
	return u.Username
}

// GetEmail return the email
func (u NewUser) GetEmail() string {
	return u.Email
}
