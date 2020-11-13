package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		register := NewRegistrator(NewFinder(db), auth.NewDefaultValidator(), db)
		auth.TestRegister(t, register, auth.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		})
		return nil
	}))
}

func TestDuplicateUsername(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := auth.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		}
		register := NewRegistrator(NewFinder(db), auth.NewDefaultValidator(), db)
		auth.TestDuplicateUsername(t, register, newUser)
		return nil
	}))
}

func TestDuplicateEmail(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := auth.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		}
		register := NewRegistrator(NewFinder(db), auth.NewDefaultValidator(), db)
		auth.TestDuplicateEmail(t, register, newUser)
		return nil
	}))
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		register := NewRegistrator(NewFinder(db), auth.NewDefaultValidator(), db)
		auth.TestEmptyPassword(t, register)
		return nil
	}))
}
