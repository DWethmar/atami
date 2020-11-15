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
}

// NewMessage model
type NewMessage struct {
	Text            string
	CreatedByUserID int
}
