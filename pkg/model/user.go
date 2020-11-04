package model

import (
	"time"
)

// User struct declaration
type User struct {
	ID        int64
	UID       string
	Username  string
	CreatedAt time.Time
}
