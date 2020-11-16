package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func generateTestMessages(size int) []message.CreateMessage {
	messages := make([]message.CreateMessage, size)
	for i := 0; i < size; i++ {
		messages[i] = message.CreateMessage{
			Text:            fmt.Sprintf("Lorum ipsum %d", i+1),
			CreatedByUserID: 1,
		}
	}
	return messages
}

func setup(db *sql.DB) (*message.Finder, []message.Message) {
	service := New(db)
	messages := make([]message.Message, 100)

	for i, newMSG := range generateTestMessages(100) {
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

func TestReadAll(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		finder, messages := setup(db)
		message.TestFindAll(t, finder, 100, messages)
		return nil
	}))
}
