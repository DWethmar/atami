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

	user *User
}

// User output
type User struct {
	ID       int
	UID      string
	Username string
}

// CreateMessage model
type CreateMessage struct {
	Text            string
	CreatedByUserID int
}
