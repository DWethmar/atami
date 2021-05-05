package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/message"
)

// findRepository reads messages from memory
type findRepository struct {
	db database.Transaction
}

// FindAll get multiple messages
func (i *findRepository) Find(limit, offset int) ([]*message.Message, error) {
	return querySelectMessages(i.db, limit, offset)
}

// FindByID get one message
func (i *findRepository) FindByUID(UID string) (*message.Message, error) {
	m, err := queryRowSelectMessageByUID(i.db, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, message.ErrCouldNotFind
		}
		return nil, err
	}
	return m, nil
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
func NewFinder(db database.Transaction) *message.Finder {
	return message.NewFinder(&findRepository{db})
}
