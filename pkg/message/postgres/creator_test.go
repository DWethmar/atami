package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newMessage := message.NewMessage{
			Text:            "wow",
			CreatedByUserID: 1,
		}
		message.TestCreator(t, NewCreator(db), newMessage)
		return nil
	}))
}
