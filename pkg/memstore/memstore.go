package memstore

import (
	"errors"
	"sync"
)

var (
	// ErrCouldNotParse error declaration
	ErrCouldNotParse = errors.New("could not parse")
)

// Memstore contains the schema.
type Memstore struct {
	user                 *UserStore
	setUserStoreState    setUserStoreState
	message              *MessageStore
	setMessageStoreState setMessageStoreState

	readMux  *sync.Mutex
	writeMux *sync.Mutex
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
func (s *Memstore) copy(kv *KvStore) (*Memstore, error) {
	store := createStore(kv)

	if users, err := s.user.All(); err == nil {
		for _, user := range users {
			store.GetUsers().Put(user.ID, user)
		}
	} else {
		return nil, err
	}

	if messages, err := s.message.All(); err == nil {
		for _, message := range messages {
			store.GetMessages().Put(message.ID, message)
		}
	} else {
		return nil, err
	}

	return store, nil
}

type txFn = func(memstore *Memstore) error

// Transaction runs transaction
func (s *Memstore) Transaction(txFn txFn) error {
	s.writeMux.Lock()
	defer s.writeMux.Unlock()

	kv := NewKvStore(&sync.Mutex{})
	copy, err := s.copy(kv)
	if err != nil {
		return err
	}

	s.readMux.Lock()
	defer s.readMux.Unlock()

	if err = txFn(copy); err != nil {
		return err
	}

	s.setUserStoreState(copy.user.GetIDs(), kv)
	s.setMessageStoreState(copy.message.GetIDs(), kv)

	return nil
}

func createStore(kv *KvStore) *Memstore {
	readMux := &sync.Mutex{}
	writeMux := &sync.Mutex{}

	userStore, setuserStoreState := NewUserStore(kv, readMux, writeMux)
	messsageStore, setMessageStoreState := NewMessageStore(kv, readMux, writeMux)

	return &Memstore{
		user:                 userStore,
		setUserStoreState:    setuserStoreState,
		message:              messsageStore,
		setMessageStoreState: setMessageStoreState,

		readMux:  readMux,
		writeMux: writeMux,
	}
}

// NewStore returns a new store
func NewStore() *Memstore {
	return createStore(NewKvStore(&sync.Mutex{}))
}
