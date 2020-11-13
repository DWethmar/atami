package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
)

var deleteMessage = fmt.Sprintf(`
DELETE FROM %s
WHERE id = $1
`, tableName)

// deleterRepository deletes user from memory
type deleterRepository struct {
	db *sql.DB
}

// Delete deletes one user
func (i deleterRepository) Delete(ID model.MessageID) error {
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
