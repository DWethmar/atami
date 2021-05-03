package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		updater := NewUpdater(db)
		test.TestUpdater(t, updater, user.UpdateUser{
			Biography: "lorum ipsum",
		})
		return nil
	}))
}
