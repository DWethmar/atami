package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/domain/message"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	db *sql.DB
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	r, err := execDeleteMessage(i.db, ID)
	if err != nil {
		return err
	}

	if a, err := r.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return message.ErrCouldNotDelete
	}

	return nil
}

// NewDeleter return a new deleter
func NewDeleter(db *sql.DB) *message.Deleter {
	return message.NewDeleter(&deleterRepository{db})
}
