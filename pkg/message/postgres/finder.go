package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/message"
)

// findRepository reads messages from memory
type findRepository struct {
	db *sql.DB
}

// FindAll get multiple messages
func (i *findRepository) Find() ([]*message.Message, error) {
	return querySelectMessages(i.db, 100, 0)
}

// FindByID get one message
func (i *findRepository) FindByID(ID int) (*message.Message, error) {
	return queryRowSelectMessageByID(i.db, ID)
}

// NewFinder return a new in memory listin reader
func NewFinder(db *sql.DB) *message.Finder {
	return message.NewFinder(&findRepository{db})
}
