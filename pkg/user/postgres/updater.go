package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

// updateRepository reads messages from memory
type updateRepository struct {
	db *sql.DB
}

// FindAll get multiple messages
func (f updateRepository) Update(updateUser user.UpdateAction) (*user.User, error) {
	return nil, nil
}

// NewUpdater return a new in memory listin repository
func NewUpdater(db *sql.DB) *user.Updater {
	return user.NewUpdater(
		&updateRepository{db},
	)
}
