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
	UpdatedAt       time.Time

	User *User
}

// User output
type User struct {
	ID       int
	UID      string
	Username string
}

// CreateMessage model
type CreateMessage struct {
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time
}
