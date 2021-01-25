package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(
			db,
		)
		msg, err := creator.Create(node.CreateRequest{
			Text:            "Lorum ipsum",
			CreatedByUserID: 1,
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, msg) {
			return errors.New("created node is nil")
		}

		deleter := NewDeleter(db)
		node.TestDelete(t, deleter, msg.ID)
		return nil
	}))
}
