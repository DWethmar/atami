package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	userMemory "github.com/dwethmar/atami/pkg/user/memory"

	userPostgres "github.com/dwethmar/atami/pkg/user/postgres"
)

// NewUserServiceMemory create a new in memory auth service
func NewUserServiceMemory() (*user.Service, *memstore.Store) {
	store := memstore.New()
	finder := userMemory.NewFinder(store)
	creator := userMemory.NewCreator(store)
	deleter := userMemory.NewDeleter(store)

	return user.NewService(*finder, *deleter, *creator), store
}

// NewUserServicePostgres create a new postgres auth service
func NewUserServicePostgres(db *sql.DB) *user.Service {
	finder := userPostgres.NewFinder(db)
	creator := userPostgres.NewCreator(db)
	deleter := userPostgres.NewDeleter(db)

	return user.NewService(*finder, *deleter, *creator)
}
