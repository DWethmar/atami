package postgres

import (
	"database/sql"

	"github.com/dwethmar/atami/pkg/message"
)

type service struct {
	message.Finder
	message.Deleter
	message.Creator
}

// New creates a new user service
func New(db *sql.DB) *message.Service {
	r := NewFinder(db)
	d := NewDeleter(db)
	c := NewCreator(db)

	return message.NewService(
		*r,
		*d,
		*c,
	)
}
