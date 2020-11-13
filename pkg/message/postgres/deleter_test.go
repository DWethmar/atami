package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		creator := NewCreator(
			db,
		)
		msg, err := creator.Create(message.NewMessage{
			Text:      "Lorum ipsum",
			CreatedBy: model.UserID(1),
		})

		if !assert.NoError(t, err) {
			return err
		}

		if !assert.NotNil(t, msg) {
			return errors.New("created message is nil")
		}

		deleter := NewDeleter(db)
		message.TestDelete(t, deleter, msg.ID)
		return nil
	}))
}