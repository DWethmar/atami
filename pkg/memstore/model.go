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

// The Message model
type Message struct {
	ID              int
	UID             string
	Text            string
	CreatedByUserID int
	CreatedAt       time.Time
}
