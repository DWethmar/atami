package domain

import (
	"database/sql"

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
}

// Store allows the mutation and reads of domain data.
type Store struct {
	Message *MessageStore
	User    *UserStore
}

// NewStore create new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Message: &MessageStore{
			Creator:   messagePostgres.NewCreator(db),
			Deleter:   messagePostgres.NewDeleter(db),
			Finder:    messagePostgres.NewFinder(db),
			Validator: message.NewValidator(),
		},
		User: &UserStore{
			Creator: userPostgres.NewCreator(db),
			Deleter: userPostgres.NewDeleter(db),
			Finder:  userPostgres.NewFinder(db),
			Updater: userPostgres.NewUpdater(db),
		},
	}
}

// NewInMemoryStore creates a store that uses inmemory storage.
func NewInMemoryStore(store *memstore.Store) *Store {
	return &Store{
		Message: &MessageStore{
			Creator:   messageMemory.NewCreator(store),
			Deleter:   messageMemory.NewDeleter(store),
			Finder:    messageMemory.NewFinder(store),
			Validator: message.NewValidator(),
		},
		User: &UserStore{
			Creator: userMemory.NewCreator(store),
			Deleter: userMemory.NewDeleter(store),
			Finder:  userMemory.NewFinder(store),
			Updater: userMemory.NewUpdater(store),
		},
	}
}
