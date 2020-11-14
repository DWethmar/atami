package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/auth"
	userMemory "github.com/dwethmar/atami/pkg/auth/memory"
	userPostgres "github.com/dwethmar/atami/pkg/auth/postgres"

	"github.com/dwethmar/atami/pkg/memstore"
)

// NewAuthServiceMemory create a new in memory auth service
func NewAuthServiceMemory() auth.Service {
	return userMemory.NewService(memstore.New())
}

// NewAuthServicePostgres create a new postgres auth service
func NewAuthServicePostgres(db *sql.DB) auth.Service {
	return userPostgres.NewService(db)
}
