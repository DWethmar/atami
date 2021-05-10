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
	CreatedBy       User
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Apply a update to the message
func (m *Message) Apply(c Update) {
	m.Text = c.Text
}

// User output
type User struct {
	ID       int
	UID      string
	Username string
}

// Create model is used toi store a message
type Create struct {
	UID             string
	Text            string
	CreatedByUserID int
}

// Update model is used to update a message
type Update struct {
	Text string
	UpdatedAt time.Time
}
