package memstore

// Store contains the schema.
type Store struct {
	user    *KvStore
	message *KvStore
}

// GetUsers return the user collection
func (s *Store) GetUsers() *KvStore {
	return s.user
}

//GetMessages return the message collection
func (s *Store) GetMessages() *KvStore {
	return s.message
}

// NewStore returns a new store
func NewStore() *Store {
	return &Store{
		user:    NewKvStore(),
		message: NewKvStore(),
	}
}
