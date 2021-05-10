package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/entity"
	"github.com/dwethmar/atami/pkg/domain/seed"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		seed.SeedUser(db, entity.NewUID(), "username", "asdasdasd", "email@email.vom", "asd", entity.Now(), entity.Now())
		updater := NewUpdater(db)
		test.TestUpdater(t, updater, user.UpdateUser{
			Biography: "lorum ipsum",
		})
		return nil
	}))
}
