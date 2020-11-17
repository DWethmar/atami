package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dwethmar/atami/pkg/message"
)

var deleteMessage = fmt.Sprintf(`
DELETE FROM %s
WHERE id = $1
`, Table)

// deleterRepository deletes user from memory
type deleterRepository struct {
	db *sql.DB
}

// Delete deletes one user
func (i deleterRepository) Delete(ID int) error {
	r, err := i.db.Exec(deleteMessage, ID)
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
