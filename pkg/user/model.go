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

	password string
}

// SetPassword sets a new password
func (u *User) SetPassword(plainPwd string, salt string) {
	u.password = hash([]byte((salt + plainPwd)))
}

// GetPassword returns the passsword
func (u *User) GetPassword() string {
	return u.password
}

// zoink returns the passsword
func (u *User) zoink() int {
	return 1
}
