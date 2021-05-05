package memstore

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	store := NewKvStore(&sync.Mutex{})

	store.Put("1", "test")

	v, ok := store.Get("1")
	assert.True(t, ok)

	value, ok := v.(string)
	assert.True(t, ok)
	assert.Equal(t, value, "test")
}

func TestDelete(t *testing.T) {
	store := NewKvStore(&sync.Mutex{})

	store.Put("1", "test uno")
	store.Put("2", "test dos")

	ok := store.Delete("1")
	assert.True(t, ok)

	_, ok = store.Get("1")
	assert.False(t, ok)
	assert.Equal(t, 1, len(store.entries))
}

func TestLen(t *testing.T) {
	store := NewKvStore(&sync.Mutex{})

	for i := 0; i < 9000; i++ {
		store.Put(strconv.Itoa(i+1), "test"+strconv.Itoa(i+1))
	}

	assert.Equal(t, 9000, len(store.entries))
}
