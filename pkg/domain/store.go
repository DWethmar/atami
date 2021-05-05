package domain

import (
	"database/sql"
	"errors"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
	messageMemory "github.com/dwethmar/atami/pkg/domain/message/memory"
	messagePostgres "github.com/dwethmar/atami/pkg/domain/message/postgres"
	userMemory "github.com/dwethmar/atami/pkg/domain/user/memory"
	userPostgres "github.com/dwethmar/atami/pkg/domain/user/postgres"

	"github.com/dwethmar/atami/pkg/memstore"

	"github.com/dwethmar/atami/pkg/domain/user"
)

// MessageStore allows the mutation and reads on message data.
type MessageStore struct {
	*message.Creator
	*message.Deleter
	*message.Finder
	*message.Validator
}

// UserStore allows the mutation and reads on user data.
type UserStore struct {
	*user.Creator
	*user.Deleter
	*user.Finder
	*user.Updater
	*user.Validator
}

// DataStore gives access to the domain models.
type DataStore struct {
	Message *MessageStore
	User    *UserStore
}

type transactionFn = func(store *DataStore) error

// Store allows the mutation and reads of domain data.
type Store struct {
	*DataStore
	execTransaction func(fn transactionFn) error
}

// Transaction creates a new transaction
func (s *Store) Transaction(fn transactionFn) error {
	if s.execTransaction == nil {
		return errors.New("store transaction unavailable")
	}
	return s.execTransaction(fn)
}

func newPostgresDataStore(db database.Transaction) *DataStore {
	var messageCreator = messagePostgres.NewCreator(db)
	var messageDeleter = messagePostgres.NewDeleter(db)
	var messageFinder = messagePostgres.NewFinder(db)
	var messageValidator = message.NewValidator()

	var userFinder = userPostgres.NewFinder(db)
	var userCreator = userPostgres.NewCreator(db, userFinder)
	var userDeleter = userPostgres.NewDeleter(db)
	var userValidator = user.NewValidator()

	return &DataStore{
		Message: &MessageStore{
			Creator:   messageCreator,
			Deleter:   messageDeleter,
			Finder:    messageFinder,
			Validator: messageValidator,
		},
		User: &UserStore{
			Creator:   userCreator,
			Deleter:   userDeleter,
			Finder:    userFinder,
			Validator: userValidator,
		},
	}
}

// NewStore create new Store
func NewStore(db *sql.DB) *Store {
	execTxFn := func(fn transactionFn) error {
		return database.WithTransaction(db, func(t database.Transaction) error {
			return fn(newPostgresDataStore(t))
		})
	}
	return &Store{
		DataStore:       newPostgresDataStore(db),
		execTransaction: execTxFn,
	}
}

// NewInMemoryStore creates a store that uses inmemory storage.
func NewInMemoryStore(store *memstore.Memstore) *Store {
	var messageCreator = messageMemory.NewCreator(store)
	var messageDeleter = messageMemory.NewDeleter(store)
	var messageFinder = messageMemory.NewFinder(store)
	var messageValidator = message.NewValidator()

	var userFinder = userMemory.NewFinder(store)
	var userCreator = userMemory.NewCreator(store, userFinder)
	var userDeleter = userMemory.NewDeleter(store)
	var userValidator = user.NewValidator()

	return &Store{
		DataStore: &DataStore{
			Message: &MessageStore{
				Creator:   messageCreator,
				Deleter:   messageDeleter,
				Finder:    messageFinder,
				Validator: messageValidator,
			},
			User: &UserStore{
				Creator:   userCreator,
				Deleter:   userDeleter,
				Finder:    userFinder,
				Validator: userValidator,
			},
		},
	}
}
