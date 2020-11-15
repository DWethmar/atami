package service

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/message"
	messageMemory "github.com/dwethmar/atami/pkg/message/memory"
	messagePostgres "github.com/dwethmar/atami/pkg/message/postgres"

	"github.com/dwethmar/atami/pkg/memstore"
)

// NewMessageServiceMemory create a new in memory message service
func NewMessageServiceMemory() *message.Service {
	return messageMemory.New(memstore.New())
}

// NewMessageServicePostgres create a new postgres message service
func NewMessageServicePostgres(db *sql.DB) *message.Service {
	return messagePostgres.New(db)
}
