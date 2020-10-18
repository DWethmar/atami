package memory

import (
	"strconv"
	"testing"

	"github.com/dwethmar/atami/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	store := NewMemstore()

	for i := 1; i <= 100; i++ {
		ok := store.Add(types.ID(i), "test"+strconv.Itoa(i))
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	assert.Equal(t, 100, len(store.List()))

	for i, e := range store.List() {
		var b = *e
		value, err := b.(string)
		assert.False(t, err)
		if !err {
			return
		}
		assert.Equal(t, "test"+strconv.Itoa(i), value)
	}
}

func TestGet(t *testing.T) {
	store := NewMemstore()

	store.Add(1, "test")

	v, ok := store.Get(1)
	assert.True(t, ok)

	var b = *v
	value, err := b.(string)
	assert.False(t, err)
	if !err {
		return
	}
	assert.Equal(t, value, "true")
}

func TestDelete(t *testing.T) {
	store := NewMemstore()

	store.Add(1, "test uno")
	store.Add(2, "test dos")

	ok := store.Delete(1)
	assert.True(t, ok)

	_, ok = store.Get(1)
	assert.False(t, ok)

	assert.Equal(t, 1, len(store.List()))
}
