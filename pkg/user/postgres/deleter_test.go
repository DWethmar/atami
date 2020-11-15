package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(
			user.NewDefaultValidator(),
			db,
		)

		u, err := creator.Create(user.CreateUser{
			Username:       "username",
			Email:          "username@test.com",
			HashedPassword: "Test1234!@#$",
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, u) {
			return errors.New("Created user is nil or ")
		}

		deleter := NewDeleter(db)
		user.TestDelete(t, deleter, u.ID)
		return nil
	}))
}
