package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func generateTestMessages(size int) []message.CreateMessageRequest {
	messages := make([]message.CreateMessageRequest, size)
	for i := 0; i < size; i++ {
		messages[i] = message.CreateMessageRequest{
			Text:            fmt.Sprintf("Lorum ipsum %d", i+1),
			CreatedByUserID: 1,
		}
	}
	return messages
}

func setup(db *sql.DB) (*message.Finder, []message.Message) {
	service := New(db)
	messages := make([]message.Message, 300)

	for i, newMSG := range generateTestMessages(300) {
		msg, err := service.Create(newMSG)
		if err != nil {
			fmt.Printf("error: %s", err)
			panic(1)
		}
		messages[i] = *msg
	}
	return NewFinder(db), messages
}

func TestReadOne(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, messages := setup(db)
		message.TestFindOne(t, finder, 10, messages[9])
		return nil
	}))
}

func TestNotFound(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, _ := setup(db)
		message.TestNotFound(t, finder)
		return nil
	}))
}

func TestFind(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, m := setup(db)

		// // Reverse items because of the order by on created_at DESC
		// for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		// 	m[i], m[j] = m[j], m[i]
		// }

		message.TestFind(t, finder, 50, m[0:51])
		return nil
	}))
}
