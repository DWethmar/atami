package user

import (
	"errors"
	"time"

	"github.com/dwethmar/atami/pkg/domain/entity"
)

var (
	// ErrUsernameAlreadyTaken error decloration
	ErrUsernameAlreadyTaken = errors.New("username is unavailable")
	// ErrEmailAlreadyTaken error decloration
	ErrEmailAlreadyTaken = errors.New("email is unavailable")
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

type UserCredentials struct {
	ID        entity.ID
	UID       entity.UID	
	Username  string
	Email     string
	Password  string
}

// Create struct declaration
type Create struct {
	UID       entity.UID
	Username  string
	Email     string
	Password  string
	Biography string	
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Update struct declaration
type Update struct {
	Biography string
	UpdatedAt time.Time
}
