package model

import (
	"time"
)

type User struct {
	ID        int64
	UID       string
	Username  string
	CreatedAt time.Time
}
