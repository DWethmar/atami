package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(user.NewDefaultValidator(), db)
		user.TestCreator(t, creator, user.CreateUser{
			Username:       "username",
			Email:          "test@test.nl",
			HashedPassword: "!Test123",
		})
		return nil
	}))
}

func TestDuplicateUsername(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := user.CreateUser{
			Username:       "username",
			Email:          "test@test.nl",
			HashedPassword: "!Test123",
		}
		creator := NewCreator(user.NewDefaultValidator(), db)
		user.TestDuplicateUsername(t, creator, newUser)
		return nil
	}))
}

func TestDuplicateEmail(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := user.CreateUser{
			Username:       "username",
			Email:          "test@test.nl",
			HashedPassword: "!Test123",
		}
		creator := NewCreator(user.NewDefaultValidator(), db)
		user.TestDuplicateEmail(t, creator, newUser)
		return nil
	}))
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(user.NewDefaultValidator(), db)
		user.TestEmptyPassword(t, creator)
		return nil
	}))
}
