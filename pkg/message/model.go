package message

import (
	"time"
)

// ID the id type that we use.
type ID int64

// The Message model
type Message struct {
	ID        ID
	Content   string
	CreatedAt time.Time
}
