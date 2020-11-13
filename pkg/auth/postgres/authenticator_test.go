package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		registrator := NewRegistrator(NewFinder(db), auth.NewDefaultValidator(), db)
		registrator.Register(auth.CreateUser{
			Username: "test",
			Email:    "test@test.com",
			Password: "TestPassw0rd0987!@#s",
		})

		authenticator := NewAuthenticator(db)
		auth.TestAuthenticate(t, authenticator, auth.Credentials{
			Email:    "test@test.com",
			Password: "TestPassw0rd0987!@#s",
		})
		return nil
	}))
}
