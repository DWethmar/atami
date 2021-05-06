package memstore

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	storeA := NewStore()
	storeA.GetMessages().Put(1, Message{
		Text: "nice",
	})

	storeA.GetMessages().Put(2, Message{
		Text: "nice",
	})

	storeA.GetUsers().Put(1, User{
		Username: "verynice",
	})

	storeA.GetUsers().Put(2, User{
		Username: "verynice",
	})

	storeB, err := storeA.copy(NewKvStore(&sync.Mutex{}))
	if assert.NoError(t, err) {
		return
	}

	messageA1, ok := storeA.GetMessages().Get(1)
	if assert.True(t, ok) {
		return
	}

	messageB1, ok := storeB.GetMessages().Get(1)
	if assert.True(t, ok) {
		return
	}

	assert.Equal(t, messageA1.Text, messageB1.Text)

	userA1, ok := storeA.GetUsers().Get(1)
	if assert.True(t, ok) {
		return
	}

	userB1, ok := storeB.GetUsers().Get(1)
	if assert.True(t, ok) {
		return
	}

	assert.Equal(t, userA1.Username, userB1.Username)
}

func TestTransaction(t *testing.T) {
	storeA := NewStore()
	storeA.GetUsers().Put(1, User{
		Username: "verynice",
	})

	storeA.GetUsers().Put(2, User{
		Username: "verynice",
	})

	storeA.Transaction(func(memstore *Memstore) error {
		memstore.GetUsers().Put(3, User{
			Username: "verynice3",
		})

		memstore.GetUsers().Put(4, User{
			Username: "verynice4",
		})
		return nil
	})

	userA3, ok := storeA.GetUsers().Get(3)
	if assert.True(t, ok) {
		return
	}
	assert.Equal(t, "verynice3", userA3.Username)
	assert.Equal(t, 4, storeA.GetUsers().Len())
}
