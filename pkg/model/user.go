package model

import (
	"strconv"
	"time"
)

// UserID type the id type used for users
type UserID int64

func (ID UserID) String() string {
	return strconv.FormatInt(int64(ID), 10)
}

// UserUID type the unique identifier for users.
type UserUID string

func (UID UserUID) String() string {
	return string(UID)
}

// User struct declaration
type User struct {
	ID        UserID
	UID       UserUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
