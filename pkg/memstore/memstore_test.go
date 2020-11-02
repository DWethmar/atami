package memstore

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	store := New()

	for i := 0; i < 100; i++ {
		ok := store.Add(strconv.Itoa(i+1), "test"+strconv.Itoa(i+1))
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	assert.Equal(t, 100, len(store.List()))

	for i, e := range store.List() {
		value, ok := e.(string)
		assert.True(t, ok)
		assert.Equal(t, "test"+strconv.Itoa(i+1), value)
	}
}

func TestGet(t *testing.T) {
	store := New()

	store.Add("1", "test")

	v, ok := store.Get("1")
	assert.True(t, ok)

	value, ok := v.(string)
	assert.True(t, ok)
	assert.Equal(t, value, "test")
}

func TestDelete(t *testing.T) {
	store := New()

	store.Add("1", "test uno")
	store.Add("2", "test dos")

	ok := store.Delete("1")
	assert.True(t, ok)

	_, ok = store.Get("1")
	assert.False(t, ok)
	assert.Equal(t, 1, len(store.List()))
}
