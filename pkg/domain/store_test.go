package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestPostgresTransaction(t *testing.T) {
	assert.NoError(t, database.WithTestDB(t, func(db *sql.DB) error {
		store := NewStore(db)
		return successfulTransaction(t, store)
	}))
}

func TestPostgresTransactionFail(t *testing.T) {
	assert.Error(t, database.WithTestDB(t, func(db *sql.DB) error {
		store := NewStore(db)
		return failTransaction(t, store)
	}))
}

func TestInMemoryTransaction(t *testing.T) {
	store := NewInMemoryStore(memstore.NewStore())
	assert.NoError(t, successfulTransaction(t, store))
}

func TestInMemoryTransactionFail(t *testing.T) {
	store := NewInMemoryStore(memstore.NewStore())
	assert.Error(t, failTransaction(t, store))
}

func successfulTransaction(t *testing.T, store *Store) error {
	msg, _ := store.Message.Create(message.CreateMessage{
		Text:            "nice",
		CreatedByUserID: 1,
	})

	return store.Transaction(func(ds *DataStore) error {
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
}

func failTransaction(t *testing.T, store *Store) error {
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

	if _, err = store.Message.FindByID(id1); assert.Error(t, err) {
		if !assert.Equal(t, err, message.ErrCouldNotFind) {
			return err
		}
	}

	if _, err = store.Message.FindByID(id2); assert.Error(t, err) {
		if !assert.Equal(t, err, message.ErrCouldNotFind) {
			return err
		}
	}

	return nil
}
