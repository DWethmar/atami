package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/test"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(
			db,
		)
		msg, err := creator.Create(message.CreateMessage{
			Text:            "Lorum ipsum",
			CreatedByUserID: 1,
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, msg) {
			return errors.New("created message is nil")
		}

		deleter := NewDeleter(db)
		test.Delete(t, deleter, msg.ID)
		return nil
	}))
}
