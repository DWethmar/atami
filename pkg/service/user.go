package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"

	userMemory "github.com/dwethmar/atami/pkg/user/memory"
	userPostgres "github.com/dwethmar/atami/pkg/user/postgres"
)

// NewUserServiceMemory create a new in memory auth service
func NewUserServiceMemory(store *memstore.Store) *user.Service {
	return userMemory.New(store)
}

// NewUserServicePostgres create a new postgres auth service
func NewUserServicePostgres(db *sql.DB) *user.Service {
	return userPostgres.New(db)
}
