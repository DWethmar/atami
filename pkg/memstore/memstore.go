package memory

import (
	"sync"

	"github.com/dwethmar/atami/pkg/types"
)

// A Memstore manages CRUD operations for entries.
type Memstore interface {
	List() []*interface{}
	Get(ID types.ID) (*interface{}, bool)
	Add(ID types.ID, value interface{}) bool
	Delete(ID types.ID) bool
}

type memstore struct {
	entries map[types.ID]interface{}
	mux     *sync.Mutex
}

// List returns all entries.
func (h *memstore) List() []*interface{} {
	var entries = make([]*interface{}, 0)
	for _, e := range h.entries {
		entries = append(entries, &e)
	}
	return entries
}

// List returns all entries.
func (h *memstore) Get(ID types.ID) (*interface{}, bool) {
	value, ok := h.entries[ID]
	return &value, ok
}

// Adds a  new entry
func (h *memstore) Add(ID types.ID, value interface{}) bool {
	h.mux.Lock()
	defer h.mux.Unlock()
	_, ok := h.entries[ID]
	if !ok {
		h.entries[ID] = &value
		return true
	}
	return false
}

func (h *memstore) Delete(ID types.ID) bool {
	h.mux.Lock()
	defer h.mux.Unlock()
	_, ok := h.entries[ID]
	if ok {
		delete(h.entries, ID)
	}
	return ok
}

// NewMemstore returns a new in memory repository.
func NewMemstore() Memstore {
	return &memstore{
		entries: make(map[types.ID]interface{}),
		mux:     &sync.Mutex{},
	}
}
