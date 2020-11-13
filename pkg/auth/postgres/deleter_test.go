package postgres

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		store := memstore.New()
		a := userRecord{
			ID:        1,
			UID:       "x",
			Email:     "test@test.nl",
			CreatedAt: time.Now(),
		}
		assert.True(t, store.Add(a.ID.String(), a))

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
		auth.TestDelete(t, deleter, a.ID)
		return nil
	}))

}
