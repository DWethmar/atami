package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/domain/user"
)

// updateRepository reads messages from memory
type updateRepository struct {
	db *sql.DB
}

// Update update user
func (f updateRepository) Update(ID int, action user.UpdateUser) (*user.User, error) {
	return queryRowUpdateUser(
		f.db,
		ID,
		action.Biography,
		action.UpdatedAt,
	)
}

// NewUpdater return a new in memory listin repository
func NewUpdater(db *sql.DB) *user.Updater {
	return user.NewUpdater(
		&updateRepository{db},
	)
}
