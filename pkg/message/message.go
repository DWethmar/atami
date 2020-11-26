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

// CreateMessageRequest model
type CreateMessageRequest struct {
	Text            string
	CreatedByUserID int
}

// CreateMessage model
type CreateMessage struct {
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time
}
