package auth

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

// RegisterUser struct declaration
type RegisterUser struct {
	Username      string
	Email         string
	PlainPassword string
}

// CreateUser struct declaration internal
type CreateUser struct {
	Username       string
	Email          string
	Salt           string
	HashedPassword string
}

// GetUsername return the username
func (u RegisterUser) GetUsername() string {
	return u.Username
}

// GetEmail return the email
func (u RegisterUser) GetEmail() string {
	return u.Email
}
