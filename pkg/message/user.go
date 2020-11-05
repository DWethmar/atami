package message

import (
	"fmt"
	"time"
)

// ID type the id type used for messages
type ID int64

func (ID ID) String() string {
	return fmt.Sprintf("%b", ID)
}

// UID type the unique identifier for referencing the message from outside.
type UID string

func (UID UID) String() string {
	return string(UID)
}

// The Message model
type Message struct {
	ID        ID
	UID       UID
	Text      string
	CreatedAt time.Time
}
