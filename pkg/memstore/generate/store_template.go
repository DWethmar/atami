package main

var storeTemplateText = `// This file was generated by robots; DO NOT EDIT.
// run: 'make generate' to regenerate this file.

package {{ .PackageName }};

import (
	"sync"
	"sort"
	"fmt"
)

type {{ .Name }}IDs = []int 
type set{{ .Name }}StoreState = func(IDs {{ .Name }}IDs, kv *KvStore) error

func generate{{ .Name }}Key(ID int) string {
	return fmt.Sprintf("{{ .Name }}_%d", ID)
}

// {{ .Name }}Store stores data in memory by key and value
type {{ .Name }}Store struct {
	kv *KvStore
	ids {{ .Name }}IDs
	readMux *sync.Mutex
	writeMux *sync.Mutex
}

// All returns all entries.
func (h *{{ .Name }}Store) All() ([]{{ .Name }}, error) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	entries := make([]{{ .Name }}, len(h.ids))

	for i, ID := range h.ids {
		record, ok := h.kv.Get(generate{{ .Name }}Key(ID))
		if ok {
			e, ok := record.({{ .Name }})
			if ok {
				entries[i] = e
			} else {
				return nil, fmt.Errorf("entry {{ .Name }} with id: %d could not be parsed", ID)
			}
		} else {
			return nil, fmt.Errorf("entry {{ .Name }} with id: %d was not found in kvstore", ID)
		}
	}

	return entries, nil
}

// Slice returns entries within the range.
func (h *{{ .Name }}Store) Slice(low, high int) ([]{{ .Name }}, error) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	entries := make([]{{ .Name }}, high-low)

	for i, ID := range h.ids[low:high] {
		if record, ok := h.kv.Get(generate{{ .Name }}Key(ID)); ok {
			if record, ok := record.({{ .Name }}); ok {
				entries[i] = record
			} else {
				return nil, fmt.Errorf("entry {{ .Name }} with id: %d could not be parsed", ID)
			}
		} else {
			return nil, fmt.Errorf("entry {{ .Name }} with id: %d was not found in kvstore", ID)
		}
	}

	return entries, nil
}

// Get a single {{ .Name }}.
func (h *{{ .Name }}Store) Get(ID int) ({{ .Name }}, bool) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	if record, ok := h.kv.Get(generate{{ .Name }}Key(ID)); ok {
		if record, ok := record.({{ .Name }}); ok {
			return record, true
		}
	}
	return {{ .Name }}{}, false
}

// Get all IDs
func (h *{{ .Name }}Store) GetIDs() {{ .Name }}IDs  {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	return h.ids
}

// Put new {{ .Name }}
func (h *{{ .Name }}Store) Put(ID int, value {{ .Name }}) bool {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	h.ids = append(h.ids, ID)
	return h.kv.Put(generate{{ .Name }}Key(ID), value)
}

// Delete a {{ .Name }}
func (h *{{ .Name }}Store) Delete(ID int) bool {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	ok := h.kv.Delete(generate{{ .Name }}Key(ID))

	if ok {
		for i, n := range h.ids {
			if n == ID {
				h.ids = append(h.ids[:i], h.ids[i+1:]...)
			}
		}
	}

	return ok
}

// Len gets number of entries
func (h *{{ .Name }}Store) Len() int {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	return len(h.ids)
}

// FromIndex gets value by index
func (h *{{ .Name }}Store) FromIndex(i int) ({{ .Name }}, bool) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	if i >= 0 && len(h.ids) > i  {
		record, ok := h.kv.Get(generate{{ .Name }}Key(h.ids[i]))
		if ok {
			if user, ok := record.({{ .Name }}); ok {
				return user, true
			}
		}
	}
	return {{ .Name }}{}, false
}

// Sort items in memory
func (h *{{ .Name }}Store) Sort(less func(i, j int) bool) {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	sort.SliceStable(h.ids, less)
}

// New{{ .Name }}Store returns a new in memory repository for {{ .Name }} records.
func New{{ .Name }}Store(kvs *KvStore, readMux *sync.Mutex, writeMux *sync.Mutex) (*{{ .Name }}Store, set{{ .Name }}StoreState) {
	store := &{{ .Name }}Store{
		kv: kvs,
		ids: make({{ .Name }}IDs, 0),
		readMux:  readMux,
		writeMux: writeMux,
	}
	return store, func(IDs {{ .Name }}IDs, kv *KvStore) error {
		store.ids = IDs
		store.kv = kv
		return nil
	}
}
`
