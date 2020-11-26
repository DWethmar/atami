package memstore

import (
	"sort"
	"sync"
)

// Store stores data in memory by key / value
type Store struct {
	entries map[string]interface{}
	order   []string
	mux     *sync.Mutex
}

// All returns all entries.
func (h *Store) All() []interface{} {
	entries := make([]interface{}, len(h.entries))

	for i, ID := range h.order {
		entries[i] = h.entries[ID]
	}

	return entries
}

// Slice returns entries within the range.
func (h *Store) Slice(low, high int) []interface{} {
	entries := make([]interface{}, high-low)

	for i, key := range h.order[low:high] {
		entries[i] = h.entries[key]
	}

	return entries
}

// Get a single value.
func (h *Store) Get(ID string) (interface{}, bool) {
	value, ok := h.entries[ID]
	return value, ok
}

// Put new value
func (h *Store) Put(key string, value interface{}) bool {
	h.mux.Lock()
	defer h.mux.Unlock()

	if _, exists := h.entries[key]; !exists {
		h.entries[key] = value
		h.order = append(h.order, key)
		return true
	}

	return false
}

// Delete a value
func (h *Store) Delete(key string) bool {
	h.mux.Lock()
	defer h.mux.Unlock()
	_, ok := h.entries[key]
	if ok {
		delete(h.entries, key)
		for i, n := range h.order {
			if n == key {
				h.order = append(h.order[:i], h.order[i+1:]...)
			}
		}
	}

	return ok
}

// Len gets number of entries
func (h *Store) Len() int {
	return len(h.order)
}

// FromIndex gets value by index
func (h *Store) FromIndex(i int) (interface{}, bool) {
	if i >= 0 && i < h.Len() {
		return h.Get(h.order[i])
	}
	return nil, false
}

// Sort items in memory
func (h *Store) Sort(less func(i, j int) bool) {
	sort.SliceStable(h.order, less)
}

// New returns a new in memory repository.
func New() *Store {
	return &Store{
		entries: make(map[string]interface{}),
		order:   make([]string, 0),
		mux:     &sync.Mutex{},
	}
}
