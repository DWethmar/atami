// This file was generated by robots; DO NOT EDIT.
// run: 'make generate' to regenerate this file.

package memstore

import (
	"fmt"
	"sort"
	"sync"
)

type MessageIDs = []int
type setMessageStoreState = func(IDs MessageIDs, kv *KvStore) error

func generateMessageKey(ID int) string {
	return fmt.Sprintf("Message_%d", ID)
}

// MessageStore stores data in memory by key and value
type MessageStore struct {
	kv       *KvStore
	ids      MessageIDs
	readMux  *sync.Mutex
	writeMux *sync.Mutex
}

// All returns all entries.
func (h *MessageStore) All() ([]Message, error) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	entries := make([]Message, len(h.ids))

	for i, ID := range h.ids {
		record, ok := h.kv.Get(generateMessageKey(ID))
		if ok {
			e, ok := record.(Message)
			if ok {
				entries[i] = e
			} else {
				return nil, fmt.Errorf("entry Message with id: %d could not be parsed", ID)
			}
		} else {
			return nil, fmt.Errorf("entry Message with id: %d was not found in kvstore", ID)
		}
	}

	return entries, nil
}

// Slice returns entries within the range.
func (h *MessageStore) Slice(low, high int) ([]Message, error) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	entries := make([]Message, high-low)

	for i, ID := range h.ids[low:high] {
		if record, ok := h.kv.Get(generateMessageKey(ID)); ok {
			if record, ok := record.(Message); ok {
				entries[i] = record
			} else {
				return nil, fmt.Errorf("entry Message with id: %d could not be parsed", ID)
			}
		} else {
			return nil, fmt.Errorf("entry Message with id: %d was not found in kvstore", ID)
		}
	}

	return entries, nil
}

// Get a single Message.
func (h *MessageStore) Get(ID int) (Message, bool) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	if record, ok := h.kv.Get(generateMessageKey(ID)); ok {
		if record, ok := record.(Message); ok {
			return record, true
		}
	}
	return Message{}, false
}

// Get all IDs
func (h *MessageStore) GetIDs() MessageIDs {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	return h.ids
}

// Put new Message
func (h *MessageStore) Put(ID int, value Message) bool {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	if h.kv.Put(generateMessageKey(ID), value) {
		h.ids = append(h.ids, ID)
	} else {
		return false
	}

	return true
}

// Delete a Message
func (h *MessageStore) Delete(ID int) bool {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	ok := h.kv.Delete(generateMessageKey(ID))

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
func (h *MessageStore) Len() int {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	return len(h.ids)
}

// FromIndex gets value by index
func (h *MessageStore) FromIndex(i int) (Message, bool) {
	h.readMux.Lock()
	defer h.readMux.Unlock()

	if i >= 0 && len(h.ids) > i {
		record, ok := h.kv.Get(generateMessageKey(h.ids[i]))
		if ok {
			if user, ok := record.(Message); ok {
				return user, true
			}
		}
	}
	return Message{}, false
}

// Sort items in memory
func (h *MessageStore) Sort(less func(i, j int) bool) {
	h.writeMux.Lock()
	defer h.writeMux.Unlock()

	sort.SliceStable(h.ids, less)
}

// NewMessageStore returns a new in memory repository for Message records.
func NewMessageStore(kvs *KvStore, readMux *sync.Mutex, writeMux *sync.Mutex) (*MessageStore, setMessageStoreState) {
	store := &MessageStore{
		kv:       kvs,
		ids:      make(MessageIDs, 0),
		readMux:  readMux,
		writeMux: writeMux,
	}
	return store, func(IDs MessageIDs, kv *KvStore) error {
		store.ids = IDs
		store.kv = kv
		return nil
	}
}
