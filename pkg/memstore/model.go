package memstore

import "time"

// User struct declaration
type User struct {
	ID        int
	UID       string
	Username  string
	Email     string
	Password  string
	Biography string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// MessageUser user related data on message
type MessageUser struct {
	ID       int
	UID      string
	Username string
}

// The Message model
type Message struct {
	ID              int
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
