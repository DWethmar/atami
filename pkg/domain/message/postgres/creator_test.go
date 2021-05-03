package postgres

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/test"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		newMessage := message.CreateMessage{
			Text:            "wow",
			CreatedByUserID: 1,
		}
		test.Create(t, NewCreator(db), newMessage)
		return nil
	}))
}
