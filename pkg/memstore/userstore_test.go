package memstore

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserList(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	for i := 0; i < 100; i++ {
		ok := store.Put(i+1, User{
			Username: "Miauw" + strconv.Itoa(i),
		})
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	assert.Equal(t, 100, store.Len())

	users, err := store.All()
	if assert.NoError(t, err) {
		for i, u := range users {
			assert.Equal(t, "Miauw"+strconv.Itoa(i), u.Username)
		}
	}
}

func TestUserSlice(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	for i := 0; i < 100; i++ {
		ok := store.Put(i+1, User{
			Username: "Miauw" + strconv.Itoa(i),
		})
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	users, err := store.Slice(10, 21)
	if assert.NoError(t, err) {
		for i, u := range users {
			assert.Equal(t, "Miauw"+strconv.Itoa(10+i), u.Username)
		}
	}
}

func TestUserGet(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	ok := store.Put(1, User{
		Username: "Miauw1",
	})

	if assert.True(t, ok) {
		v, ok := store.Get(1)
		assert.True(t, ok)
		assert.Equal(t, "Miauw1", v.Username)
	}
}

func TestUserDelete(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	store.Put(1, User{
		Username: "Miauw1",
	})

	store.Put(2, User{
		Username: "Miauw2",
	})

	ok := store.Delete(1)
	assert.True(t, ok)

	_, ok = store.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 1, store.Len())
}

func TestUserLen(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	for i := 0; i < 9000; i++ {
		store.Put(i+1, User{
			Username: "Miauw" + strconv.Itoa(i),
		})
	}
	assert.Equal(t, 9000, store.Len())
}

func TestUserFromIndex(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	for i := 0; i < 9000; i++ {
		store.Put(i+1, User{
			Username: "Miauw" + strconv.Itoa(i),
		})
	}

	if value, ok := store.FromIndex(0); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, "Miauw0", value.Username)
	}

	if value, ok := store.FromIndex(100); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, "Miauw100", value.Username)
	}

	if value, ok := store.FromIndex(8999); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, "Miauw8999", value.Username)
	}
}

// Sort items in memory
func TestUserSort(t *testing.T) {
	store := NewUserStore(&sync.Mutex{})

	store.Put(3, User{
		Username:  "Miauw3",
		CreatedAt: time.Now().UTC().Add(time.Duration(3000)),
	})

	store.Put(2, User{
		Username:  "Miauw2",
		CreatedAt: time.Now().UTC().Add(time.Duration(2000)),
	})

	store.Put(1, User{
		Username:  "Miauw1",
		CreatedAt: time.Now().UTC().Add(time.Duration(1000)),
	})

	store.Sort(func(i, j int) bool {
		a, ok := store.FromIndex(i)
		if !ok {
			return false
		}

		b, ok := store.FromIndex(j)
		if !ok {
			return false
		}

		return a.Username < b.Username
	})

	if value, ok := store.FromIndex(0); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		assert.Equal(t, "Miauw1", value.Username)
	}

	if value, ok := store.FromIndex(1); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		assert.Equal(t, "Miauw2", value.Username)
	}

	if value, ok := store.FromIndex(2); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		assert.Equal(t, "Miauw3", value.Username)
	}

	if _, ok := store.FromIndex(3); ok {
		assert.Fail(t, "there should be no fourth value!!")
	}
}
