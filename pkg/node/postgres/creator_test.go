package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newNode := node.CreateRequest{
			Text:            "wow",
			CreatedByUserID: 1,
		}
		node.TestCreator(t, NewCreator(db), newNode)
		return nil
	}))
}
