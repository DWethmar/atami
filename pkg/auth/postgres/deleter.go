package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/auth"
)

var deleteUser = `
DELETE FROM public.user
WHERE id = $1
`

// deleterRepository deletes user from memory
type deleterRepository struct {
	db *sql.DB
}

// Delete deletes one user
func (i deleterRepository) Delete(ID auth.ID) error {
	r, err := i.db.Exec(deleteUser, ID)
	if err != nil {
		return err
	}

	if a, err := r.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return auth.ErrCouldNotDelete
	}

	return nil
}

// NewDeleter return a new in deleter repo
func NewDeleter(db *sql.DB) *auth.Deleter {
	return auth.NewDeleter(&deleterRepository{db})
}
