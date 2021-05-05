package domain

import (
	"database/sql"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestPostgresTransaction(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		ds := NewStore(db)

		msg, _ := ds.Message.Create(message.CreateMessage{
			Text:            "nice",
			CreatedByUserID: 1,
		})

		err := ds.Transaction(func(ds *DataStore) error {
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
		ds := NewStore(db)

		var id1 int
		var id2 int

		err := ds.Transaction(func(ds *DataStore) error {
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

			panic(1)
		})

		_, err = ds.Message.FindByID(id1)
		assert.Equal(t, err, message.ErrCouldNotFind)

		_, err = ds.Message.FindByID(id2)
		assert.Equal(t, err, message.ErrCouldNotFind)

		return nil
	}))
}
