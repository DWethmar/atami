package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	db *sql.DB
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	r, err := i.db.Exec(deleteUser, ID)
	if err != nil {
		return err
	}

	if a, err := r.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return user.ErrCouldNotDelete
	}

	return nil
}

// NewDeleter return a new in deleter repo
func NewDeleter(db *sql.DB) *user.Deleter {
	return user.NewDeleter(&deleterRepository{db})
}
