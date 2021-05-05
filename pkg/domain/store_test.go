package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestPostgresTransaction(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		store := NewStore(db)

		msg, _ := store.Message.Create(message.CreateMessage{
			Text:            "nice",
			CreatedByUserID: 1,
		})

		err := store.Transaction(func(ds *DataStore) error {
			usr, err := ds.User.Create(user.CreateUser{
				Username: "mrtest",
				Password: "askjldash3kljd&*&sdsK<LJLIHJ",
				Email:    "testtest@test.nl",
			})
			if err != nil {
				return err
			}

			_, err = ds.Message.Create(message.CreateMessage{
				Text:            "nice",
				CreatedByUserID: usr.ID,
			})
			if err != nil {
				return err
			}

			if err := ds.Message.Delete(msg.ID); err != nil {
				return err
			}

			return nil
		})

		assert.NoError(t, err)

		return nil
	}))
}

func TestPostgresTransactionFail(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		store := NewStore(db)

		var id1 int
		var id2 int

		err := store.Transaction(func(ds *DataStore) error {
			msg1, _ := ds.Message.Create(message.CreateMessage{
				Text:            "nice",
				CreatedByUserID: 1,
			})
			id1 = msg1.ID

			msg2, _ := ds.Message.Create(message.CreateMessage{
				Text:            "nice",
				CreatedByUserID: 1,
			})
			id2 = msg2.ID

			return errors.New("something went wrong")
		})

		if id1 == 0 || id2 == 0 {
			assert.Fail(t, fmt.Sprintf("one of the ids is 0: id1: %d, %d", id1, id2))
		}

		_, err = store.Message.FindByID(id1)
		assert.Equal(t, err, message.ErrCouldNotFind)

		_, err = store.Message.FindByID(id2)
		assert.Equal(t, err, message.ErrCouldNotFind)

		return nil
	}))
}
