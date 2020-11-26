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
func (i *findRepository) Find(limit, offset int) ([]*message.Message, error) {
	return querySelectMessages(i.db, limit, offset)
}

// FindByID get one message
func (i *findRepository) FindByID(ID int) (*message.Message, error) {
	m, err := queryRowSelectMessageByID(i.db, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, message.ErrCouldNotFind
		}
		return nil, err
	}
	return m, nil
}

// NewFinder return a new in memory listin reader
func NewFinder(db *sql.DB) *message.Finder {
	return message.NewFinder(&findRepository{db})
}
