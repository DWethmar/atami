package memstore

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	store := New()

	for i := 0; i < 100; i++ {
		ok := store.Put(strconv.Itoa(i+1), "test"+strconv.Itoa(i+1))
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	assert.Equal(t, 100, len(store.All()))

	for i, e := range store.All() {
		value, ok := e.(string)
		assert.True(t, ok)
		assert.Equal(t, "test"+strconv.Itoa(i+1), value)
	}
}

func TestSlice(t *testing.T) {
	store := New()

	for i := 0; i < 100; i++ {
		ok := store.Put(strconv.Itoa(i), "test"+strconv.Itoa(i))
		assert.True(t, ok)
		if !ok {
			return
		}
	}

	for i, e := range store.Slice(10, 21) {
		value, ok := e.(string)
		if assert.True(t, ok) {
			assert.Equal(t, "test"+strconv.Itoa(i+10), value)
		}
	}
}

func TestGet(t *testing.T) {
	store := New()

	store.Put("1", "test")

	v, ok := store.Get("1")
	assert.True(t, ok)

	value, ok := v.(string)
	assert.True(t, ok)
	assert.Equal(t, value, "test")
}

func TestDelete(t *testing.T) {
	store := New()

	store.Put("1", "test uno")
	store.Put("2", "test dos")

	ok := store.Delete("1")
	assert.True(t, ok)

	_, ok = store.Get("1")
	assert.False(t, ok)
	assert.Equal(t, 1, len(store.All()))
}

func TestLen(t *testing.T) {
	store := New()
	for i := 0; i < 9000; i++ {
		store.Put(strconv.Itoa(i+1), "test"+strconv.Itoa(i+1))
	}
	assert.Equal(t, 9000, store.Len())
}

func TestFromIndex(t *testing.T) {
	store := New()
	for i := 0; i < 9000; i++ {
		store.Put(strconv.Itoa(i+1), "test"+strconv.Itoa(i))
	}

	if value, ok := store.FromIndex(0); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, value, "test0")
	}

	if value, ok := store.FromIndex(100); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, value, "test100")
	}

	if value, ok := store.FromIndex(8999); assert.True(t, ok) && assert.NotNil(t, value) {
		assert.Equal(t, value, "test8999")
	}
}

// Sort items in memory
func TestSort(t *testing.T) {
	store := New()

	type testObj struct {
		Number int
		Text   string
	}

	store.Put("1", testObj{
		Number: 99,
	})
	store.Put("2", testObj{
		Number: 44,
	})
	store.Put("3", testObj{
		Number: 0,
	})

	store.Sort(func(i, j int) bool {
		a, ok := store.FromIndex(i)
		if !ok {
			return false
		}
		aa, ok := a.(testObj)

		b, ok := store.FromIndex(j)
		if !ok {
			return false
		}
		bb, ok := b.(testObj)

		return aa.Number > bb.Number
	})

	if value, ok := store.FromIndex(0); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		first, ok := value.(testObj)
		assert.True(t, ok)
		assert.Equal(t, 99, first.Number, "first value")
	}

	if value, ok := store.FromIndex(1); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		second, ok := value.(testObj)
		assert.True(t, ok)
		assert.Equal(t, 44, second.Number, "second value")
	}

	if value, ok := store.FromIndex(2); ok {
		assert.True(t, ok)
		assert.NotNil(t, value)
		third, ok := value.(testObj)
		assert.True(t, ok)
		assert.Equal(t, 0, third.Number, "third value")
	}

	if _, ok := store.FromIndex(3); ok {
		assert.Fail(t, "there should be no fourth value")
	}
}
