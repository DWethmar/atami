package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(db)
		test.TestCreator(t, creator, user.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		})
		return nil
	}))
}

func TestDuplicateUsername(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := user.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		}
		creator := NewCreator(db)
		test.TestDuplicateUsername(t, creator, newUser)
		return nil
	}))
}

func TestDuplicateEmail(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newUser := user.CreateUser{
			Username: "username",
			Email:    "test@test.nl",
			Password: "!Test123",
		}
		creator := NewCreator(db)
		test.TestDuplicateEmail(t, creator, newUser)
		return nil
	}))
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(db)
		test.TestEmptyPassword(t, creator)
		return nil
	}))
}
