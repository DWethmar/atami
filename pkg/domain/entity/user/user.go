package user

import (
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

// User struct declaration
type User struct {
	ID        entity.ID
	UID       entity.UID
	Username  string
	Email     string
	Password  string
	Biography string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create struct declaration
type Create struct {
	UID       entity.UID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Update struct declaration
type Update struct {
	Biography string
	UpdatedAt time.Time
}
