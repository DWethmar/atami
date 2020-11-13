package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		registrator := NewRegistrator(
			NewFinder(db),
			auth.NewDefaultValidator(),
			db,
		)

		user, err := registrator.Register(auth.CreateUser{
			Username: "username",
			Email:    "username@test.com",
			Password: "Test1234!@#$",
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, user) {
			return errors.New("Created user is nil or ")
		}

		deleter := NewDeleter(db)
		auth.TestDelete(t, deleter, user.ID)
		return nil
	}))
}
