package memstore

import (
	"sync"
)

// KvStore stores data in memory by key / value
type KvStore struct {
	entries map[string]interface{}
	mux     *sync.Mutex
}

// Get a single value.
func (h *KvStore) Get(key string) (interface{}, bool) {
	h.mux.Lock()
	defer h.mux.Unlock()

	value, ok := h.entries[key]
	return value, ok
}

// Put new value
func (h *KvStore) Put(key string, value interface{}) bool {
	h.mux.Lock()
	defer h.mux.Unlock()

	if _, exists := h.entries[key]; !exists {
		h.entries[key] = value
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
	}

	return ok
}

// NewKvStore returns a new in memory repository.
func NewKvStore(mux *sync.Mutex) *KvStore {
	return &KvStore{
		entries: make(map[string]interface{}),
		mux:     mux,
	}
}
