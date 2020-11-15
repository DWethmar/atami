package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	userMemory "github.com/dwethmar/atami/pkg/user/memory"

	userPostgres "github.com/dwethmar/atami/pkg/user/postgres"
)

// NewAuthServiceMemory create a new in memory auth service
func NewAuthServiceMemory(store *memstore.Store) *auth.Service {
	finder := userMemory.NewFinder(store)
	creator := userMemory.NewCreator(store)

	return auth.NewService(
		*auth.NewAuthenticator(finder),
		*auth.NewRegistrator(creator, finder),
	)
}

// NewAuthServicePostgres create a new postgres auth service
func NewAuthServicePostgres(db *sql.DB) *auth.Service {
	finder := userPostgres.NewFinder(db)
	creator := userPostgres.NewCreator(db)

	return auth.NewService(
		*auth.NewAuthenticator(finder),
		*auth.NewRegistrator(creator, finder),
	)
}
