package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/domain/message"
	messageMemory "github.com/dwethmar/atami/pkg/domain/message/memory"
	messagePostgres "github.com/dwethmar/atami/pkg/domain/message/postgres"

	"github.com/dwethmar/atami/pkg/memstore"
)

// NewMessageServiceMemory create a new in memory message service
func NewMessageServiceMemory(store *memstore.Store) *message.Service {
	return messageMemory.New(store)
}

// NewMessageServicePostgres create a new postgres message service
func NewMessageServicePostgres(db *sql.DB) *message.Service {
	return messagePostgres.New(db)
}
