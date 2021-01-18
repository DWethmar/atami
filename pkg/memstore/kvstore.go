package memstore

import (
	"sort"
	"sync"
)

// KvStore stores data in memory by key / value
type KvStore struct {
	entries map[string]interface{}
	order   []string
	mux     *sync.Mutex
}

// All returns all entries.
func (h *KvStore) All() []interface{} {
	h.mux.Lock()
	defer h.mux.Unlock()
	entries := make([]interface{}, len(h.entries))

	for i, ID := range h.order {
		entries[i] = h.entries[ID]
	}

	return entries
}

// Slice returns entries within the range.
func (h *KvStore) Slice(low, high int) []interface{} {
	h.mux.Lock()
	defer h.mux.Unlock()
	entries := make([]interface{}, high-low)

	for i, key := range h.order[low:high] {
		entries[i] = h.entries[key]
	}

	return entries
}

// Get a single value.
func (h *KvStore) Get(ID string) (interface{}, bool) {
	value, ok := h.entries[ID]
	return value, ok
}

// Put new value
func (h *KvStore) Put(key string, value interface{}) bool {
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
func (h *KvStore) Delete(key string) bool {
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
func (h *KvStore) Len() int {
	return len(h.order)
}

// FromIndex gets value by index
func (h *KvStore) FromIndex(i int) (interface{}, bool) {
	if i >= 0 && i < h.Len() {
		return h.Get(h.order[i])
	}
	return nil, false
}

// Sort items in memory
func (h *KvStore) Sort(less func(i, j int) bool) {
	sort.SliceStable(h.order, less)
}

// NewKvStore returns a new in memory repository.
func NewKvStore() *KvStore {
	return &KvStore{
		entries: make(map[string]interface{}),
		order:   make([]string, 0),
		mux:     &sync.Mutex{},
	}
}
