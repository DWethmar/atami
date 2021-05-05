package postgres

import (
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain/user"
)

// deleterRepository deletes user from memory
type deleterRepository struct {
	db database.Transaction
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	r, err := execDeleteUser(i.db, ID)
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
func NewDeleter(db database.Transaction) *user.Deleter {
	return user.NewDeleter(&deleterRepository{db})
}
