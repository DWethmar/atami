package model

import (
	"strconv"
	"time"
)

// MessageID type the id type used for messages
type MessageID int64

func (ID MessageID) String() string {
	return strconv.FormatInt(int64(ID), 10)
}

// MessageUID type the unique identifier for referencing the message from outside.
type MessageUID string

func (UID MessageUID) String() string {
	return string(UID)
}

// The Message model
type Message struct {
	ID        MessageID
	UID       MessageUID
	Text      string
	CreatedAt time.Time
}
