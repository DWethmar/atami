package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(db)

		u, err := creator.Create(user.CreateUser{
			Username: "username",
			Email:    "username@test.com",
			Password: "Test1234!@#$",
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, u) {
			return errors.New("Created user is nil or ")
		}

		deleter := NewDeleter(db)
		test.TestDelete(t, deleter, u.ID)
		return nil
	}))
}
