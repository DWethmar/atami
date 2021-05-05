package memstore

import (
	"errors"
	"sync"
)

var (
	// ErrCouldNotParse error declaration
	ErrCouldNotParse = errors.New("Could not parse")
)

// Memstore contains the schema.
type Memstore struct {
	user    *UserStore
	message *MessageStore

	mux *sync.Mutex
}

// GetUsers return the user collection
func (s *Memstore) GetUsers() *UserStore {
	return s.user
}

//GetMessages return the message collection
func (s *Memstore) GetMessages() *MessageStore {
	return s.message
}

// Copy memstore data
func (s *Memstore) Copy() *Memstore {
	return nil
}

// NewStore returns a new store
func NewStore() *Memstore {
	mux := &sync.Mutex{}

	return &Memstore{
		user:    NewUserStore(mux),
		message: NewMessageStore(mux),
		mux:     mux,
	}
}
