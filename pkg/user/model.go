package user

import (
	"fmt"
	"time"
)

// ID type the id type used for users
type ID int64

func (ID ID) String() string {
	return fmt.Sprintf("%b", ID)
}

// UID type the unique identifier for users.
type UID string

func (UID UID) String() string {
	return string(UID)
}

// User struct declaration
type User struct {
	ID        ID
	UID       UID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
}
