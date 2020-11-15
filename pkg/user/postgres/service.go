package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/user"
)

// New creates a new user service
func New(db *sql.DB) *user.Service {
	var validator = user.NewDefaultValidator()

	f := NewFinder(db)
	d := NewDeleter(db)
	r := NewCreator(validator, db)

	return user.NewService(
		*f,
		*d,
		*r,
		*validator,
	)
}
