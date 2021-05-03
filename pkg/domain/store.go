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

type messageStore struct {
	*message.Creator
	*message.Deleter
	*message.Finder
}

type userStore struct {
	*user.Creator
	*user.Deleter
	*user.Finder
	*user.Updater
}

type Store struct {
	Message messageStore
	User userStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Message: messageStore{
			Creator: messagePostgres.NewCreator(db),
			Deleter: messagePostgres.NewDeleter(db),
			Finder: messagePostgres.NewFinder(db),
		},
		User: userStore{
			Creator: userPostgres.NewCreator(db),
			Deleter: userPostgres.NewDeleter(db),
			Finder: userPostgres.NewFinder(db),
			Updater: userPostgres.NewUpdater(db),	
		},
	}
}

func NewInMemoryStore(store *memstore.Store) *Store {
	return &Store{
		Message: messageStore{
			Creator: messageMemory.NewCreator(store),
			Deleter: messageMemory.NewDeleter(store),
			Finder: messageMemory.NewFinder(store),
		},
		User: userStore{
			Creator: userMemory.NewCreator(store),
			Deleter: userMemory.NewDeleter(store),
			Finder: userMemory.NewFinder(store),
			Updater: userMemory.NewUpdater(store),	
		},
	}
}