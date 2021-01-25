package memstore

import "errors"

var (
	// ErrCouldNotParse error declaration
	ErrCouldNotParse = errors.New("Could not parse")
)

// Store contains the schema.
type Store struct {
	user    *UserStore
	message *MessageStore
	node    *NodeStore
}

// GetUsers return the user collection
func (s *Store) GetUsers() *UserStore {
	return s.user
}

//GetMessages return the message collection
func (s *Store) GetMessages() *MessageStore {
	return s.message
}

// GetNodes return the node collection
func (s *Store) GetNodes() *NodeStore {
	return s.node
}

// NewStore returns a new store
func NewStore() *Store {
	return &Store{
		user:    NewUserStore(),
		message: NewMessageStore(),
		node:    NewNodeStore(),
	}
}
