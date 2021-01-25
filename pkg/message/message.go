package message

import (
	"time"
)

// The Message model
type Message struct {
	ID              int
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time

	User *User
}

// User output
type User struct {
	ID       int
	UID      string
	Username string
}

// CreateRequest model
type CreateRequest struct {
	Text            string
	CreatedByUserID int
}

// CreateAction model
type CreateAction struct {
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time
}
