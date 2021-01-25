package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

// New creates a new user service
func New(db *sql.DB) *user.Service {
	f := NewFinder(db)
	d := NewDeleter(db)
	r := NewCreator(db)
	u := NewUpdater(db)

	return user.NewService(
		*f,
		*d,
		*r,
		*u,
	)
}
