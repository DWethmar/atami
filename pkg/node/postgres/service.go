package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/node"
)

// New creates a new user service
func New(db *sql.DB) *node.Service {
	r := NewFinder(db)
	d := NewDeleter(db)
	c := NewCreator(db)

	return node.NewService(
		*r,
		*d,
		*c,
	)
}
